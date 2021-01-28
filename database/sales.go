package database

import (
	"fmt"

	"github.com/yy-c00/hotel-demo/model"
)

type salesManager struct{}

func NewSalesManager() model.SalesManager {
	return salesManager{}
}

func (sm salesManager) NewSale(sale *model.Sale) error {
	if sale.Room == (model.Room{}) && sale.Credential == (model.Credential{}) {
		return model.ErrFailedSale
	}

	tx, err := Connection().Begin()
	if err != nil {
		return err
	}

	lastInsert, err := tx.Exec("INSERT INTO sales() VALUES()")
	if err != nil {
		return err
	}

	lastSaleID, err := lastInsert.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	sale.ID = uint(lastSaleID)
	for _, product := range sale.Products {
		_, err = tx.Exec("INSERT INTO sold(idproducts, quantity, idsales) VALUES(?, ?, ?)", product.ID, product.Quantity, sale.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if sale.Room.ID > 0 {
		_, err = tx.Exec("INSERT INTO room_sales VALUES(?, ?)", sale.Room.ID, sale.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if sale.Credential != (model.Credential{}) {
		if sale.Credential.Root {
			_, err = tx.Exec("INSERT INTO root_sales VALUES (?, ?)", sale.Credential.ID, sale.ID)
		} else {
			_, err = tx.Exec("INSERT INTO employee_sales VALUES (?, ?)", sale.Credential.ID, sale.ID)
		}

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	var count uint
	err = tx.QueryRow("SELECT COUNT(*) FROM products WHERE quantity < 0").Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}

	if count > 0 {
		tx.Rollback()
		return model.ErrFailedSale
	}

	return tx.Commit()
}

func (sm salesManager) ConfirmSale(sale model.Sale) error {
	var table string

	if sale.Credential.Root {
		table = "root"
	} else {
		table = "employee"
	}

	_, err := Connection().Exec(fmt.Sprintf("INSERT INTO %s_sales VALUES(?, ?)", table), sale.Credential.ID, sale.ID)
	return err
}

func (sm salesManager) CancelSale(idSale uint) error {
	_, err := Connection().Exec("DELETE FROM sales WHERE idsales = ?", idSale)
	return err
}

func (sm salesManager) GetAllPendingSales() ([]model.Sale, error) {
	var sales []model.Sale

	salesRows, _ := Connection().Query("SELECT idsales FROM sales WHERE status = FALSE")
	defer salesRows.Close()

	for salesRows.Next() {
		var idSale uint

		err := salesRows.Scan(&idSale)
		if err != nil {
			return []model.Sale{}, err
		}

		newSale, err := sm.GetSaleById(idSale)
		if err != nil {
			return []model.Sale{}, err
		}

		sales = append(sales, newSale)
	}

	return sales, nil
}

func (sm salesManager) GetSaleById(idSale uint) (model.Sale, error) {
	newSale := model.Sale{ID: idSale}

	err := Connection().QueryRow("SELECT idsales, status, time_stamp FROM sales WHERE idsales = ?", idSale).Scan(&newSale.ID, &newSale.Status, &newSale.Date)
	if err != nil {
		return model.Sale{}, err
	}

	productsRows, err := Connection().Query("CALL GET_PRODUCTS_FROM_SALE(?)", newSale.ID)
	if err != nil {
		return model.Sale{}, err
	}
	defer productsRows.Close()

	products := []model.Product{}
	for productsRows.Next() {
		newProduct := model.Product{}

		err := productsRows.Scan(&newProduct.ID, &newProduct.Code, &newProduct.Name, &newProduct.Price, &newProduct.Quantity)
		if err != nil {
			return model.Sale{}, err
		}

		products = append(products, newProduct)
	}

	err = Connection().QueryRow("CALL SEARCH_ROOM_SALE(?)", newSale.ID).Scan(&newSale.Room.ID)
	newSale.Products = products

	return newSale, nil
}
