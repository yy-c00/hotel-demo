package router

import (
	"github.com/yy-c00/hotel-demo/authorization"
	"github.com/yy-c00/hotel-demo/database"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yy-c00/hotel-demo/model"
)

type userAccessRouter struct {
	model.UserAccessManager
}

func NewUserAccessRouter() model.UserAccessRouter {
	return userAccessRouter{database.NewUserAccessManager()}
}

func (a userAccessRouter) UserIsLogged(c echo.Context) error {
	credential, ok := c.Get("user").(model.Credential)
	if !ok {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "no se ha encontrado la llave"})
	}

	err := a.UserAccessManager.UserIsLogged(credential)
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, credential)
}

func (a userAccessRouter) ValidateUser(c echo.Context) error {
	newAccess := model.Access{}

	err := c.Bind(&newAccess)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	credential, err := a.UserAccessManager.ValidateUser(newAccess)
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"message": model.ErrAccessDenied.Error()})
	}

	token, err := authorization.GenerateToken(&credential, 24 * time.Hour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (a userAccessRouter) CreateNewUser(c echo.Context) error {
	user := model.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	credential, err := a.UserAccessManager.CreateNewUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, credential)
}

func (a userAccessRouter) AvailableUsers(c echo.Context) error {
	credentials, err := a.UserAccessManager.AvailableUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, credentials)
}