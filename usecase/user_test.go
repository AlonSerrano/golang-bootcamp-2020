package usecase

import (
	"github.com/AlonSerrano/GolangBootcamp/models"
	"github.com/AlonSerrano/GolangBootcamp/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUp(t *testing.T) {
	u := models.User{
		Id:             "",
		FirstName:      "Test",
		SecondName:     "Tes",
		LastName:       "Te",
		SecondLastName: "T",
		Birthdate:      "22/11/1991",
		Email:          "email@mail.com",
		Phone:          "5512345678",
		Password:       "12345",
	}
	collection := service.UseUserTable()
	assert.Condition(t, func() bool {
		_, err := SignUp(u, collection)
		if err != nil {
			return false
		}
		return true
	}, "Add User")
}

func TestSignIn(t *testing.T) {
	u := models.SignInReq{
		Email:    "email@mail.com",
		Password: "12345",
	}
	collection := service.UseUserTable()
	assert.Condition(t, func() bool {
		_, err := SignIn(u, collection)
		if err != nil {
			return false
		}
		return true
	}, "SignIn User")
}
