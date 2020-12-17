package usecase

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AlonSerrano/GolangBootcamp/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/encoding/charmap"
)

var isoStateAbbreviation = [32]string{"MX-AGU", "MX-BCN", "MX-BCS", "MX-CAM", "MX-CHP", "MX-CHH", "MX-COA", "MX-COL", "MX-CMX", "MX-DUR", "MX-GUA", "MX-GRO", "MX-HID", "MX-JAL", "MX-MEX", "MX-MIC", "MX-MOR", "MX-NAY", "MX-NLE", "MX-OAX", "MX-PUE", "MX-QUE", "MX-ROO", "MX-SLP", "MX-SIN", "MX-SON", "MX-TAB", "MX-TAM", "MX-TLA", "MX-VER", "MX-YUC", "MX-ZAC"}

// UpdateRecords the function obtains the zip codes then deletes the table and finishes inserting the new data
func UpdateRecords(collection *mongo.Collection, c echo.Context) models.SignUpRes {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	zipCodes, zipCodesModel := getCSVCodes()
	dropZipCodes(collection)
	saveZipCodes(zipCodes, collection)
	return models.SignUpRes{
		Message: "Have been inserted " + strconv.Itoa(len(zipCodesModel)) + " postal codes updated by " + email,
	}
}

// SearchZipCodes Search the zipcodes database the neighborhood by zip code
func SearchZipCodes(cp string, collection *mongo.Collection) []models.ZipCodeCSV {
	findOpts := options.Find()
	var results []models.ZipCodeCSV
	filter := bson.D{{"codigoPostal", cp}}
	cur, err := collection.Find(context.TODO(), filter, findOpts)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var s models.ZipCodeCSV
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, s)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}

func dropZipCodes(collection *mongo.Collection) bool {
	collection.Drop(context.TODO())
	logrus.Info("Dropped table")
	return true
}

func saveZipCodes(zip []interface{}, collection *mongo.Collection) *mongo.InsertManyResult {
	insertManyResult, err := collection.InsertMany(context.TODO(), zip)
	if err != nil {
		log.Fatal(err)
	}
	return insertManyResult
}

func getCSVCodes() ([]interface{}, []models.ZipCodeCSV) {
	createZipCodesCSVFile()
	csvFile, err := os.Open("zipcodes.csv")
	if err != nil {
		logrus.Error("Couldn't open the csv file", err)
	}
	zipCodes, zipCodesModel := csvToMap(csvFile)
	return zipCodes, zipCodesModel
}

func createZipCodesCSVFile() {
	url := viper.GetString("ZipCodesDB")
	response, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
	}
	defer response.Body.Close()
	stream := charmap.ISO8859_1.NewDecoder().Reader(response.Body)
	csvFile, err := os.Create("zipcodes.csv")
	if err != nil {
		logrus.Error("failed creating file: ", err)
	}
	defer csvFile.Close()
	_, err = io.Copy(csvFile, stream)
	if err != nil {
		logrus.Error(err)
	}
}

func csvToMap(reader io.Reader) ([]interface{}, []models.ZipCodeCSV) {
	var zipCodes []interface{}
	var zipCodesModel []models.ZipCodeCSV
	r := csv.NewReader(reader)
	r.Comma = '|'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	if _, err := r.Read(); err != nil {
		panic(err)
	}
	header, err := r.Read()
	if err != nil {
		panic(err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		dict := map[string]string{}
		for i := range header {
			dict[header[i]] = record[i]
		}
		u := strings.Replace(uuid.New().String(), "-", "", -1)
		isoCode, err := strconv.Atoi(dict["c_estado"])
		if err != nil {
			continue
		}
		parsingZipCode := models.ZipCodeCSV{
			Id:           u,
			CodigoPostal: dict["d_codigo"],
			Estado:       dict["d_estado"],
			EstadoISO:    isoStateAbbreviation[isoCode-1],
			Municipio:    dict["D_mnpio"],
			Ciudad:       dict["d_ciudad"],
			Barrio:       dict["d_asenta"],
		}
		zipCodes = append(zipCodes, parsingZipCode)
		zipCodesModel = append(zipCodesModel, parsingZipCode)
	}
	return zipCodes, zipCodesModel
}
