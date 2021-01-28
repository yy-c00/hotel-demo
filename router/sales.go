package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/olahol/melody"
	"github.com/yy-c00/hotel-demo/database"
	"github.com/yy-c00/hotel-demo/model"
)

type salesRouter struct {
	model.SalesManager
	channel *melody.Melody
}

func NewSalesRouter() model.SalesRouter {
	channel := melody.New()

	channel.HandleConnect(connect)

	return salesRouter{database.NewSalesManager(), channel}
}

func (sales salesRouter) NewSale(c echo.Context) error {
	var sale model.Sale

	err := c.Bind(&sale)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = sales.SalesManager.NewSale(&sale)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	msg := model.Message{
		Type: model.NewSale,
		Data: sale,
	}

	err = c.JSON(http.StatusCreated, echo.Map{"message": "ok"})
	if err != nil {
		return err
	}

	bytes, _ := json.Marshal(msg)
	sales.channel.Broadcast(bytes)

	return nil
}

func (sales salesRouter) ConfirmSale(c echo.Context) error {
	credential, ok := c.Get("user").(model.Credential)
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
	}


	err = sales.SalesManager.ConfirmSale(model.Sale{ID: uint(id), Credential: credential})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err})
	}

	confirmedSale, err := sales.SalesManager.GetSaleById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	msg := model.Message{
		Type: model.ConfirmedSale,
		Data: confirmedSale,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err})
	}
	sales.channel.Broadcast(bytes)

	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (sales salesRouter) CancelSale(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	canceledSale, err := sales.SalesManager.GetSaleById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	err = sales.SalesManager.CancelSale(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	msg := model.Message{
		Type: model.CanceledSale,
		Data: canceledSale,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	sales.channel.Broadcast(bytes)
	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (sales salesRouter) GetAllPendingSales(c echo.Context) error {
	allSales, err := sales.SalesManager.GetAllPendingSales()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, allSales)
}

func (sales salesRouter) GetSaleById(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	sale, err := sales.SalesManager.GetSaleById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, sale)
}

func (sales salesRouter) Channel(c echo.Context) error {
	sales.channel.HandleRequest(c.Response(), c.Request())
	return nil
}