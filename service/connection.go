package service

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnDB() *mongo.Client {
	host := viper.GetString("DBAddress")
	port := viper.GetInt64("DBPort")
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connections
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Congratulations, you're already connected to MongoDB!")
	return client
}

// UseZipCodeTable This method is to establish the connection with the mongo database and select the table ZipCodes
func UseZipCodeTable() *mongo.Collection {
	client := getConnDB()
	return client.Database("bootcamp").Collection("ZipCodes")
}

// UseUserTable This method is to establish the connection with the mongo database and select the table User
func UseUserTable() *mongo.Collection {
	client := getConnDB()
	return client.Database("bootcamp").Collection("User")
}
