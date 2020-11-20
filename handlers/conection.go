package handlers

import (
	"github.com/AlonSerrano/GolangBootcamp/service"
	"github.com/AlonSerrano/GolangBootcamp/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type PerissonConnectorHandler struct {
	client *mongo.Client
}

// NewPerissonConnectorHandler Initializes the connection with the database and returns the handler
func NewPerissonConnectorHandler() *PerissonConnectorHandler {
	return &PerissonConnectorHandler{client: service.GetConnDB()}
}

// HandleSearchZipCodes handler that invokes the postal code search function
func (hc *PerissonConnectorHandler) HandleSearchZipCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, usecase.SearchZipCodes(c.Param("zipCode"), hc.client))
}

// HandlePopulateZipCodes handler that invokes the postal code populate function
func (hc *PerissonConnectorHandler) HandlePopulateZipCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, usecase.GetAndSave(hc.client))
}
