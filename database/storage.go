package database

import (
	"github.com/yy-c00/hotel-demo/model"
	"strconv"
)

type productStorage struct{}

//NewStorage returns a model.Storage instance
func NewStorage() model.Storage {
	return productStorage{}
}

func (ps productStorage) AddProduct(id uint, product *model.Product) error {
	if product.Name == "" {
		return model.ErrEmptyField
	}

	insertProducts, err := Connection().Prepare("INSERT INTO products(code, name, price, idcategory) VALUES(?, ?, ?, ?)")
	defer insertProducts.Close()
	if err != nil {
		return err
	}

	lastInsert, err := insertProducts.Exec(product.Code, product.Name, product.Price, id)
	if err != nil {
		return err
	}

	lastId, err := lastInsert.LastInsertId()

	product.ID = uint(lastId)

	return err
}

func (ps productStorage) SetProduct(product model.Product) error {
	if product.Name == "" {
		return model.ErrEmptyField
	}

	_, err := Connection().Exec("UPDATE products SET code = ?, name = ?, price = ? WHERE idproducts = ?", product.Code, product.Name, product.Price, product.ID)
	return err
}

func (ps productStorage) DeleteProduct(ID uint) error {
	_, err := Connection().Exec("DELETE FROM products WHERE idproducts = ?", ID)
	return err
}

func (ps productStorage) AddCategory(category *model.Category) error {
	if category.Name == "" {
		return model.ErrEmptyField
	}

	tx, err := Connection().Begin()
	if err != nil {
		return err
	}

	lastInsert, err := tx.Exec("INSERT INTO category(name) VALUES(?)", category.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	lastId, err := lastInsert.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, product := range category.Products {
		err = ps.AddProduct(uint(lastId), &product)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	category.ID = uint(lastId)

	return tx.Commit()
}

func (ps productStorage) AddCategories(categories []model.Category) error {
	for _, category := range categories {
		err := ps.AddCategory(&category)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ps productStorage) DeleteCategory(id uint) error {
	_, err := Connection().Exec("DELETE FROM category WHERE idcategory = ?", id)
	return err
}

func (ps productStorage) GetAll() ([]model.Category, error) {
	//Procedure
	var categories []model.Category

	allCategories, err := Connection().Query("SELECT * FROM category")

	if err != nil {
		return nil, err
	}
	defer allCategories.Close()

	for allCategories.Next() {
		newCategory := model.Category{}

		allCategories.Scan(&newCategory.ID, &newCategory.Name)

		allProducts, err := Connection().Query("SELECT * FROM products WHERE idcategory = ?", newCategory.ID)
		if err != nil {
			return nil, err
		}

		for allProducts.Next() {
			newProduct := model.Product{}

			err = allProducts.Scan(&newProduct.ID, &newProduct.Code, &newProduct.Name, &newProduct.Price, &newProduct.Quantity, new(uint))
			if err != nil {
				return nil, err
			}

			newCategory.Products = append(newCategory.Products, newProduct)
		}

		allProducts.Close()
		categories = append(categories, newCategory)
	}

	return categories, nil
}

func (ps productStorage) GetCategoryByID(id uint) (model.Category, error) {
	var (
		products []model.Product
		category = model.Category{ID: id}
	)

	err := Connection().QueryRow("SELECT name FROM category WHERE idcategory = ?", id).Scan(&category.Name)
	if err != nil {
		return category, err
	}

	results, err := Connection().Query("SELECT idproducts, code, name, price, quantity FROM products WHERE idcategory = ?", id)
	if err != nil {
		return category, err
	}
	defer results.Close()

	for results.Next() {
		newProduct := model.Product{}

		err = results.Scan(&newProduct.ID, &newProduct.Code, &newProduct.Name, &newProduct.Price, &newProduct.Quantity)
		if err != nil {
			return category, err
		}

		products = append(products, newProduct)
	}

	category.Products = products

	return category, nil
}

func (ps productStorage) GetProductByID(id uint) (model.Product, error) {
	product := model.Product{}

	result := Connection().QueryRow("SELECT * FROM idproducts, code, name, price, quantity WHERE idproducts = ?", id)

	err := result.Scan(&product.ID, &product.Code, &product.Name, &product.Price, &product.Quantity)

	return product, err
}

func (ps productStorage) Log(credential model.Credential, product model.Product) error {
	if !credential.Root {
		return model.ErrAccessDenied
	}

	tx, err := Connection().Begin()
	if err != nil {
		return err
	}

	insertIntoLog, err := tx.Prepare("INSERT INTO log(idroot) VALUES(?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer insertIntoLog.Close()

	lastInsert, err := insertIntoLog.Exec(credential.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	insertInto, err := tx.Prepare("INSERT INTO log_products(idlog, quantity, idproducts) VALUES(?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer insertInto.Close()

	lastId, err := lastInsert.LastInsertId()
	if err != nil {
		return err
	}

	_, err = insertInto.Exec(lastId, product.Quantity, product.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (ps productStorage) Search(str string) ([]model.Product, error){
	var products []model.Product

	id, err := strconv.Atoi(str)
	if err == nil {
		product, err := ps.GetProductByID(uint(id))
		if err == nil {
			return []model.Product{product}, err
		}
	}

	results, err := Connection().Query("SELECT idproducts, code, name, price, quantity FROM products WHERE code LIKE ? OR name LIKE ?", "%" + str + "%", "%" + str + "%")
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		newProduct := model.Product{}

		err = results.Scan(&newProduct.ID, &newProduct.Code, &newProduct.Name, &newProduct.Price, &newProduct.Quantity)
		if err != nil {
			return nil, err
		}

		products = append(products, newProduct)
	}

	return products, nil
}