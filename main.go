package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yy-c00/hotel-demo/authorization"
	"github.com/yy-c00/hotel-demo/database"
	"github.com/yy-c00/hotel-demo/router"
	"log"
	"os"
)

const defaultPort = "80"

func main() {
	server := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := authorization.LoadCertificate("./certificates/privateFile.rsa")
	if err != nil {
		log.Fatal(err)
	}

	if database.Connection().Ping() != nil {
		log.Fatal(err)
	}

	router.SetMiddlewares(server)
	router.SetRoutes(server, router.New())

	err = server.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
