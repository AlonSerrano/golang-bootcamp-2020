package main

import (
	_ "github.com/AlonSerrano/GolangBootcamp/docs"
	"github.com/AlonSerrano/GolangBootcamp/handlers"
	"github.com/AlonSerrano/GolangBootcamp/middleware"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Infof("no default config, file not found")
	}
}
func init() {
	Init()
}

// @title App
// @description Get the neighborhood by ZipCodes of Mexico
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()
	v1 := e.Group("/api/v1")

	address := v1.Group("/address")
	{
		hz := handlers.NewZipCodesConnectorHandler()
		address.GET("/populate", hz.HandlePopulateZipCodes, middleware.IsLoggedIn, middleware.IsAdmin)
		address.GET("/search/:zipCode", hz.HandleSearchZipCodes)
	}
	user := v1.Group("/user")
	{
		ha := handlers.NewAuthHandler()
		user.POST("/signIn", ha.SignIn)
		user.POST("/signUp", ha.SignUp)
	}
	serverAddress := viper.GetString("ServerAddress")
	logrus.Info("server started at ", serverAddress)

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.Logger.Fatal(e.Start(serverAddress))
}
