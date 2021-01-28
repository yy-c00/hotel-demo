package database

import (
	"github.com/yy-c00/hotel-api/model"
	"reflect"
	"testing"
)

var (
	sales = NewSalesManager()
	newSale        = model.Sale{
		Credential: model.Credential{ID: 1},
		Room: model.Room{ID: 1},
	}
)

func TestSalesManager_NewSale(t *testing.T) {
	err := sales.NewSale(&newSale)
	if err != nil {
		t.Errorf("Ocurrs an unexpected error (%v)", err)
	}
}

func TestSalesManager_GetSaleById(t *testing.T) {
	sale, err := sales.GetSaleById(newSale.ID)
	if err != nil {
		t.Errorf("Ocurrs an unexpected error (%v)", err)
	}

	if !reflect.DeepEqual(sale, newSale) {
		t.Errorf("Expected: %v Got: %v", sale, newSale)
	}
}

func TestSalesManager_ConfirmSale(t *testing.T) {
	newSale.Credential = model.Credential{ID: 1}

	err := sales.ConfirmSale(newSale)
	if err != nil {
		t.Errorf("Ocurrs an unexpected error (%v)", err)
	}
}

func TestSalesManager_CancelSale(t *testing.T) {
	newSale.Credential = model.Credential{ID: 1}

	err := sales.CancelSale(newSale.ID)
	if err != nil {
		t.Errorf("Ocurrs an unexpected error (%v)", err)
	}
}
