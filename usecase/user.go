package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/AlonSerrano/GolangBootcamp/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//SignUp The logic to save to de database the request of a new user
func SignUp(user models.User, collection *mongo.Collection) (map[string]string, error) {
	pwd := []byte(user.Password)
	u := strings.Replace(uuid.New().String(), "-", "", -1)
	userM := models.User{
		Id:             u,
		FirstName:      user.FirstName,
		SecondName:     user.SecondName,
		LastName:       user.LastName,
		SecondLastName: user.SecondLastName,
		Birthdate:      user.Birthdate,
		Email:          user.Email,
		Phone:          user.Phone,
		Password:       hashAndSalt(pwd),
	}
	_, err := collection.InsertOne(context.TODO(), userM)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return map[string]string{
		"Status":  "200",
		"Message": "Registered user successfully",
	}, nil
}

//SignIn The logic to retrieve token of an existing user
func SignIn(signInReq models.SignInReq, collection *mongo.Collection) (*models.TokenInfoResp, error) {
	filter := bson.D{{
		Key:   "email",
		Value: signInReq.Email,
	}}
	var dbRes *models.User
	err := collection.FindOne(context.TODO(), filter).Decode(&dbRes)
	if err != nil {
		return nil, err
	}
	pwd := []byte(signInReq.Password)
	if comparePasswords(dbRes.Password, pwd) {
		tokens, err := generateTokenPair(dbRes)
		return tokens, err
	}
	return nil, echo.ErrUnauthorized
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logrus.Error(err)
	}
	return string(hash)
}

func generateTokenPair(user *models.User) (*models.TokenInfoResp, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["firstName"] = user.FirstName
	claims["lastName"] = user.LastName
	claims["sub"] = 1
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	secret := "secret"
	secretByte := []byte(secret)
	t, err := token.SignedString(secretByte)
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString(secretByte)
	if err != nil {
		return nil, err
	}
	return &models.TokenInfoResp{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return true
}
