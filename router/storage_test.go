package router

import (
	"github.com/yy-c00/hotel-demo/model"
	"net/http"
	"testing"
)

var (
	storageRouter = NewStorageRouter()
	category      = model.Category{
		Name:     "Some category",
		Products: []model.Product{{Name: "Some product"}},
	}
	configs []RequestConfig
)

func init() {
	configs = []RequestConfig{
		{
			"AddCategory",
			"",
			&category,
			http.MethodPost,
			storageRouter.AddCategory,
			http.StatusCreated,
		},
		{
			"AddProduct",
			"id=1",
			&category.Products[0],
			http.MethodPost,
			storageRouter.AddProduct,
			http.StatusCreated,
		},
		{
			"AddCategories",
			"",
			&[]model.Category{category},
			http.MethodPost,
			storageRouter.AddCategories,
			http.StatusCreated,
		},
	}
}

func TestStorage(t *testing.T) {
	ConfigTest(configs, t)
}
