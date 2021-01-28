package router

import (
	"github.com/labstack/echo/v4"
	"github.com/yy-c00/hotel-demo/database"
	"github.com/yy-c00/hotel-demo/model"
	"net/http"
	"strconv"
)

type historyRouter struct {
	model.History
}

func NewHistoryRouter() model.HistoryRouter {
	return historyRouter{database.NewHistory()}
}

func (h historyRouter) GetLogsByRange(c echo.Context) error {
	left, err := strconv.Atoi(c.QueryParam("left"))
	if err != nil {
		return c.JSON(http.StatusOK, echo.ErrBadRequest)
	}

	right, err := strconv.Atoi(c.QueryParam("right"))
	if err != nil {
		return c.JSON(http.StatusOK, echo.ErrBadRequest)
	}

	sales, err := h.History.GetLogsByRange(uint(left), uint(right))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, sales)
}

func (h historyRouter) GetSaleLogById(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	sale, err := h.History.GetSaleLogById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, sale)
}

func (h historyRouter) SearchByRange(c echo.Context) error {
	left, err := strconv.Atoi(c.QueryParam("left"))
	if err != nil {
		return c.JSON(http.StatusOK, echo.ErrBadRequest)
	}

	right, err := strconv.Atoi(c.QueryParam("right"))
	if err != nil {
		return c.JSON(http.StatusOK, echo.ErrBadRequest)
	}

	sales, err := h.History.SearchByRange(c.QueryParam("search"), uint(left), uint(right))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, sales)
}