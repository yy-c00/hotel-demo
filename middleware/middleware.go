package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yy-c00/hotel-api/authorization"
	"github.com/yy-c00/hotel-api/model"
	"net/http"
)

//ToUser provides a echo.MiddlewareFunc to validate user permissions
func ToUser() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(configToUser(authorization.Secret()))
}

//ToRoom provides a echo.MiddlewareFunc to validate room permissions
func ToRoom() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(configToRoom(authorization.Secret()))
}

func onlyRoot(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		credential, ok := c.Get("user").(model.Credential)
		if !ok || !credential.Root {
			return c.JSON(http.StatusForbidden, echo.ErrForbidden)
		}

		c.Set("user", nil)
		c.Set("root", credential)
		return handlerFunc(c)
	}
}

func toRoot(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return ToUser()(onlyRoot(handlerFunc))
}

//ToRoot provides a echo.MiddlewareFunc to validate root permissions
func ToRoot() echo.MiddlewareFunc {
	return toRoot
}

//ToRoom provides a echo.MiddlewareFunc to validate if user or room is logged
func ToAnyOne(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	configUser, configRoom := configToUser(authorization.Secret()), configToRoom(authorization.Secret())

	configUser.Skipper = func(c echo.Context) bool {
		_, err := c.Cookie("room")
		return err == nil
	}

	configRoom.Skipper = func(c echo.Context) bool {
		_, err := c.Cookie("user")
		return err == nil
	}

	toUser, toRoom := middleware.JWTWithConfig(configUser), middleware.JWTWithConfig(configRoom)

	return toUser(toRoom(handlerFunc))
}

//configUser.Skipper = func(c echo.Context) bool {
//	_, err := c.Cookie("user")
//	return err == nil
//}
//
//configRoom.Skipper = func(c echo.Context) bool {
//	_, err := c.Cookie("user")
//	return err == nil
//}