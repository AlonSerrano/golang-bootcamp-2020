package main

import (
	"log"

	"github.com/AlonSerrano/GolangBootcamp/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	v1 := e.Group("/api/v1")
	h := handlers.NewPerissonConnectorHandler()
	address := v1.Group("/address")
	{
		address.GET("/populate", h.HandlePopulateZipCodes)
		address.GET("/search/:zipCode", h.HandleSearchZipCodes)
	}
	serverAddress := "localhost:8080"
	log.Printf("server started at %s\n", serverAddress)
	e.Logger.Fatal(e.Start(serverAddress))
}
