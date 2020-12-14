package handlers

import (
	"net/http"

	"github.com/AlonSerrano/GolangBootcamp/service"
	"github.com/AlonSerrano/GolangBootcamp/usecase"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppConnectorHandler this structure is to define the base of the functions
type AppConnectorHandler struct {
	client *mongo.Collection
}

// NewZipCodesConnectorHandler Initializes the connection with the database and returns the handler
func NewZipCodesConnectorHandler() *AppConnectorHandler {
	return &AppConnectorHandler{client: service.UseZipCodeTable()}
}

// HandleSearchZipCodes handler that invokes the postal code search function
// @Summary Search by zipCode
// @Description When invoking this method, the neighborhoods belonging to a postal code will be obtained
// @Tags address
// @Accept json
// @Produce json
// @Param zipCode path int true "Zip Code"
// @Success 200 {array} models.ZipCodeCSV
// @Router /address/search/{zipCode} [get]
func (hc *AppConnectorHandler) HandleSearchZipCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, usecase.SearchZipCodes(c.Param("zipCode"), hc.client))
}

// HandlePopulateZipCodes handler that invokes the postal code populate function
// @Summary Populate the DB from Internet
// @Description When invoking this method, the postal codes will be obtained in a .csv file to proceed with the elimination of the current table in mongo and finish with the insertion with the new file
// @Tags address
// @Produce json
// @Success 200 {object} models.SignUpRes
// @Router /address/populate [get]
// @Security ApiKeyAuth
func (hc *AppConnectorHandler) HandlePopulateZipCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, usecase.UpdateRecords(hc.client, c))
}
