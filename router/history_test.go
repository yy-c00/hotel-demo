package router

import (
	"net/http"
	"testing"
)

var (
	history = NewHistoryRouter()
	historyTable = []RequestConfig {
		{
			"GetLogsByRange",
			"left=0&right=5",
			nil,
			http.MethodGet,
			history.GetLogsByRange,
			http.StatusOK,
		},
		{
			"GetSaleLogById",
			"id=1",
			nil,
			http.MethodGet,
			history.GetSaleLogById,
			http.StatusOK,
		},
		{
			"SearchByRange",
			"",
			nil,
			http.MethodGet,
			history.SearchByRange,
			http.StatusOK,
		},
	}
)

func TestHistoryRouter(t *testing.T) {
	ConfigTest(historyTable, t)
}