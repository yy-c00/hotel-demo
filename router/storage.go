package router

import (
	"net/http"
	"strconv"

	"github.com/yy-c00/hotel-demo/database"

	"github.com/labstack/echo/v4"
	"github.com/yy-c00/hotel-demo/model"
)

type storage struct {
	model.Storage
}

func NewStorageRouter() model.StorageRouter {
	return storage{database.NewStorage()}
}

func (store storage) AddProduct(c echo.Context) error {
	newProduct := model.Product{}

	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = store.Storage.AddProduct(uint(id), &newProduct)
	if err == model.ErrEmptyField {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, newProduct)
}

func (store storage) SetProduct(c echo.Context) error {
	newProduct := model.Product{}

	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = store.Storage.SetProduct(newProduct)
	if err == model.ErrEmptyField {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (store storage) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = store.Storage.DeleteProduct(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (store storage) AddCategory(c echo.Context) error {
	var newCategory model.Category

	err := c.Bind(&newCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = store.Storage.AddCategory(&newCategory)
	if err == model.ErrEmptyField {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, newCategory)
}

func (store storage) AddCategories(c echo.Context) error {
	var newProducts []model.Category

	err := c.Bind(&newProducts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = store.Storage.AddCategories(newProducts)
	if err == model.ErrEmptyField {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "ok"})
}

func (store storage) DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = store.Storage.DeleteCategory(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (store storage) GetAll(c echo.Context) error {
	allProducts, err := store.Storage.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, allProducts)
}

func (store storage) GetProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{"message": err.Error()})
	}

	product, err := store.Storage.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}

func (store storage) GetCategoryByID(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{"message": err.Error()})
	}

	category, err := store.Storage.GetCategoryByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, category)
}

func (store storage) Log(c echo.Context) error {
	newProduct := model.Product{}

	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	credential, ok := c.Get("root").(model.Credential)
	if !ok {
		return c.JSON(http.StatusForbidden, echo.ErrForbidden)
	}

	err = store.Storage.Log(credential, newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (store storage) Search(c echo.Context) error {
	str := c.QueryParam("q")

	products, err := store.Storage.Search(str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err})
	}

	return c.JSON(http.StatusOK, products)
}

func (store storage) Channel(c echo.Context) error {
	panic(model.ErrFunctionIsNotDefined)
}