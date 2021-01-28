package database

import (
	"fmt"
	"github.com/yy-c00/hotel-api/model"
	"reflect"
	"testing"
)

var (
	storage = NewStorage()
	categories []model.Category
)

func init() {
	categories = *new([]model.Category)

	for x := 1; x < 5; x++ {
		category := model.Category{Name: fmt.Sprint("Category", x)}
		products := *new([]model.Product)

		for y := 1; y < 5; y++ {
			product := model.Product{Name: fmt.Sprint("Product", y)}
			products = append(products, product)
		}

		categories = append(categories, category)
	}
}

func TestProductStorage_AddCategory(t *testing.T) {
	err := storage.AddCategory(&model.Category{Name: "Category 0"})
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}
}

func TestProductStorage_AddProduct(t *testing.T) {
	err := storage.AddProduct(1, &model.Product{Name: "Product 0"})
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}
}

func TestProductStorage_DeleteProduct(t *testing.T) {
	err := storage.DeleteProduct(1)
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}
}

func TestProductStorage_DeleteCategory(t *testing.T) {
	err := storage.DeleteCategory(1)
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}
}

func TestProductStorage_AddCategories(t *testing.T) {
	err := storage.AddCategories(categories)
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}
}

func TestProductStorage_GetAll(t *testing.T) {
	arr, err := storage.GetAll()
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}

	for i := 0; i < len(arr); i++ {
		if !reflect.DeepEqual(arr[i], categories[i]) {
			t.Errorf("Expected: %v\nGot: %v", categories[i], arr[i])
		}
	}
}