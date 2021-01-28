package model

import "errors"

//ErrEmptyField returns when a field that can no be empty is empty
var ErrEmptyField = errors.New("the field value can not be empty")

//Category category
type Category struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

//Product product
type Product struct {
	ID       uint    `json:"id"`
	Code     uint    `json:"code"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

//Storage interface implemented by all structs to manage the storage
type Storage interface {
	//AddProduct is used to add a product in some category
	AddProduct(uint, *Product) error
	//SetProduct is used to update a existing product from the storage
	SetProduct(Product) error
	//DeleteProduct is used to delete a product from the storage
	DeleteProduct(uint) error

	//AddCategory is used to add a category to storage
	AddCategory(*Category) error
	//AddCategories is used to add various categories to storage
	AddCategories([]Category) error
	//DeleteCategory is used to delete an existing category
	DeleteCategory(uint) error

	//GetAll is used to get all existing categories
	GetAll() ([]Category, error)
	//GetProductByID is used to get a product by id
	GetProductByID(uint) (Product, error)
	//GetCategoryByID is used to get a category by id
	GetCategoryByID(uint) (Category, error)

	//Log is used to add and log a new quantity of some product
	Log(Credential, Product) error
	//Search is used to search into history
	Search(string) ([]Product, error)
}
