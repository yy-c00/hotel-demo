package router

import (
	"github.com/labstack/echo/v4"
	"github.com/yy-c00/hotel-demo/authorization"
	"github.com/yy-c00/hotel-demo/database"
	"github.com/yy-c00/hotel-demo/model"
	"net/http"
	"time"
)


type roomAccessRouter struct {
	model.RoomAccessManager
}

func NewRoomAccessRouter() model.RoomAccessRouter {
	return roomAccessRouter{database.NewRoomAccessManager()}
}

func (r roomAccessRouter) RoomIsLogged(c echo.Context) error {
	room, ok := c.Get("room").(model.Room)
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": echo.ErrBadRequest.Error()})
	}

	err := r.RoomAccessManager.ValidateRoom(room)
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, room)
}

func (r roomAccessRouter) ValidateRoom(c echo.Context) error {
	room := model.Room{}

	err := c.Bind(&room)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	err = r.RoomAccessManager.ValidateRoom(room)
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
	}

	token, err := authorization.GenerateToken(&room, 24 * time.Hour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (r roomAccessRouter) CreateNewRoom(c echo.Context) error {
	newRoom := model.Room{}

	err := c.Bind(&newRoom)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = r.RoomAccessManager.CreateNewRoom(newRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func (r roomAccessRouter) AvailableRooms(c echo.Context) error {
	rooms, err := r.RoomAccessManager.AvailableRooms()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, rooms)
}