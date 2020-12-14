package handlers

import (
	"github.com/AlonSerrano/GolangBootcamp/service"
	"github.com/AlonSerrano/GolangBootcamp/usecase"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// PerissonConnectorHandler this structure is to define the base of the functions
type PerissonConnectorHandler struct {
	client *mongo.Collection
}

// NewPerissonConnectorHandler Initializes the connection with the database and returns the handler
func NewPerissonConnectorHandler() *PerissonConnectorHandler {
	return &PerissonConnectorHandler{client: service.UseZipCodeTable()}
}

// HandleSearchZipCodes handler that invokes the postal code search function
func (hc *PerissonConnectorHandler) HandleSearchZipCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, usecase.SearchZipCodes(c.Param("zipCode"), hc.client))
}

// HandlePopulateZipCodes handler that invokes the postal code populate function
func (hc *PerissonConnectorHandler) HandlePopulateZipCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, usecase.GetAndSave(hc.client))
}
