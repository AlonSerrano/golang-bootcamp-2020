package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlonSerrano/GolangBootcamp/models"
	"github.com/AlonSerrano/GolangBootcamp/service"
	"github.com/AlonSerrano/GolangBootcamp/usecase"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type authHandler struct {
	client *mongo.Collection
}

//NewAuthHandler NewAuthHandler init
func NewAuthHandler() *authHandler {
	return &authHandler{client: service.UseUserTable()}
}

// SignIn handler that invokes signIn function
// @Summary SignIn a user into the app
// @Description By sending email and password it will return a token object
// @Tags user
// @Param userParams body models.SignInReq true "User params"
// @Produce json
// @Success 200 {object} models.TokenInfoResp
// @Router /user/signIn [post]
func (ah *authHandler) SignIn(c echo.Context) error {
	jsonMap := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	} else {
		email := jsonMap["email"]
		password := jsonMap["password"]
		res, err := usecase.SignIn(models.SignInReq{
			Email:    email,
			Password: password,
		}, ah.client)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	}
}

// SignUp handler that invokes the singUp function
// @Summary SignUp Register a user in the system
// @Description By sending params it will register a user
// @Tags user
// @Produce json
// @Success 200 {object} models.TokenInfoResp
// @Accept  json
// @Param user body models.User true "User Info"
// @Router /user/signUp [post]
func (ah *authHandler) SignUp(c echo.Context) error {
	jsonMap := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		res, err := usecase.SignUp(models.User{
			Id:             "",
			FirstName:      jsonMap["firstName"],
			SecondName:     jsonMap["secondName"],
			LastName:       jsonMap["lastName"],
			SecondLastName: jsonMap["secondLastName"],
			Birthdate:      jsonMap["birthdate"],
			Email:          jsonMap["email"],
			Phone:          jsonMap["phone"],
			Password:       jsonMap["password"],
		}, ah.client)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}
