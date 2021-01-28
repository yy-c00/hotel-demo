package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yy-c00/hotel-demo/authorization"
	"github.com/yy-c00/hotel-demo/model"
)

func successUser(c echo.Context) {
	token := c.Get("token").(*jwt.Token)
	claim := token.Claims.(*authorization.CustomClaim)

	credential := model.Credential{}
	err := json.Unmarshal(claim.Data, &credential)
	if err == nil {
		c.Set("user", credential)
		c.Set("token", nil)
	}
}

func successRoom(c echo.Context) {
	token := c.Get("token").(*jwt.Token)
	claim := token.Claims.(*authorization.CustomClaim)

	room := model.Room{}
	err := json.Unmarshal(claim.Data, &room)
	if err == nil {
		c.Set("room", room)
		c.Set("token", nil)
	}

}

func configToUser(secret []byte) middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey:     secret,
		SigningMethod:  jwt.SigningMethodHS256.Name,
		TokenLookup:    "cookie:user",
		ContextKey:     "token",
		Claims:         &authorization.CustomClaim{},
		AuthScheme:     "",
		SuccessHandler: successUser,
	}
}

func configToRoom(secret []byte) middleware.JWTConfig {
	config := configToUser(secret)
	config.TokenLookup = "cookie:room"
	config.SuccessHandler = successRoom
	return config
}