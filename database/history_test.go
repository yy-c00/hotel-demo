package database

import (
	"github.com/yy-c00/hotel-api/model"
	"reflect"
	"testing"
)

var (
	historyManager model.History
	expected = *new([]model.Sale)
)

type historyMock struct {
	sales []model.Sale
}

func (m historyMock) GetLogsByRange(left uint, right uint) ([]model.Sale, error) {
	return m.sales[left:right], nil
}

func (m historyMock) GetSaleLogById(id uint) (model.Sale, error) {
	return m.sales[id - 1], nil
}

func (m historyMock) SearchByRange(str string, left uint, right uint) ([]model.Sale, error) {
	panic("implement me")
}

func init() {
	for i := uint(1); i <= 20; i++ {
		expected = append(expected, model.Sale{ID: i})
	}
	historyManager = historyMock{expected[:]}
}

func TestHistory_GetLogsByRange(t *testing.T) {
	logs, err := historyManager.GetLogsByRange(0, 4)
	if err != nil {
		t.Errorf("Occurs an unexpected error (%v)", err)
	}

	for x := uint(0); x < uint(len(logs)); x++ {
		want, err := historyManager.GetLogsByRange(x, x + 3)
		if err != nil {
			t.Errorf("Occurs an unexpected error (%v)", err)
		}

		if !reflect.DeepEqual(expected[x:x + 3], want) {
			t.Errorf("Expected: %v Got: %v", expected[x:x+3], want)
		}
	}
}

func TestHistory_GetSaleLogById(t *testing.T) {
	for i := 0; i < 10; i++ {
		sale, err := historyManager.GetSaleLogById(uint(i + 1))
		if err != nil {
			t.Errorf("Occurs an unexpected error (%v)", err)
		}
		if !reflect.DeepEqual(expected[i], sale) {
			t.Errorf("Expected: %v Got: %v", expected[i], sale)
		}
	}
}

func TestHistory_SearchByRange(t *testing.T) {

}