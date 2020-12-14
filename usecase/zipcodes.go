package usecase

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AlonSerrano/GolangBootcamp/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/encoding/charmap"
)

var rows = []map[string]string{
	{"isoCode": "MX-AGU"},
	{"isoCode": "MX-BCN"},
	{"isoCode": "MX-BCN"},
	{"isoCode": "MX-BCS"},
	{"isoCode": "MX-CAM"},
	{"isoCode": "MX-COA"},
	{"isoCode": "MX-COL"},
	{"isoCode": "MX-CHP"},
	{"isoCode": "MX-CHH"},
	{"isoCode": "MX-CMX"},
	{"isoCode": "MX-DUR"},
	{"isoCode": "MX-GUA"},
	{"isoCode": "MX-GRO"},
	{"isoCode": "MX-HID"},
	{"isoCode": "MX-JAL"},
	{"isoCode": "MX-MEX"},
	{"isoCode": "MX-MIC"},
	{"isoCode": "MX-MOR"},
	{"isoCode": "MX-NAY"},
	{"isoCode": "MX-NLE"},
	{"isoCode": "MX-OAX"},
	{"isoCode": "MX-PUE"},
	{"isoCode": "MX-QUE"},
	{"isoCode": "MX-ROO"},
	{"isoCode": "MX-SLP"},
	{"isoCode": "MX-SIN"},
	{"isoCode": "MX-SON"},
	{"isoCode": "MX-TAB"},
	{"isoCode": "MX-TAM"},
	{"isoCode": "MX-TLA"},
	{"isoCode": "MX-VER"},
	{"isoCode": "MX-YUC"},
	{"isoCode": "MX-ZAC"},
}

// GetAndSave the function obtains the zip codes then deletes the table and finishes inserting the new data
func GetAndSave(collection *mongo.Collection) *mongo.InsertManyResult {
	zipCodes, zipCodesModel := getCSVCodes()
	fmt.Println("Have been inserted %i postal codes", len(zipCodesModel))
	dropZipCodes(collection)
	return saveZipCodes(zipCodes, collection)
}

func dropZipCodes(collection *mongo.Collection) bool {
	collection.Drop(context.TODO())
	fmt.Printf("Droppped table")
	return true
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

func saveZipCodes(zip []interface{}, collection *mongo.Collection) *mongo.InsertManyResult {
	insertManyResult, err := collection.InsertMany(context.TODO(), zip)
	if err != nil {
		log.Fatal(err)
	}
	return insertManyResult
}

func getCSVCodes() ([]interface{}, []models.ZipCodeCSV) {
	getZipCodesCSVFile()
	csvFile, err := os.Open("zipcodes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	zipCodes, zipCodesModel := csvToMap(csvFile)
	return zipCodes, zipCodesModel
}

func getZipCodesCSVFile() {
	url := "https://www.correosdemexico.gob.mx/datosabiertos/cp/cpdescarga.txt"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	stream := charmap.ISO8859_1.NewDecoder().Reader(response.Body)
	csvFile, err := os.Create("zipcodes.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	_, err = io.Copy(csvFile, stream)
	if err != nil {
		log.Fatal(err)
	}
}

func csvToMap(reader io.Reader) ([]interface{}, []models.ZipCodeCSV) {
	var zipCodes []interface{}
	var zipCodesModel []models.ZipCodeCSV
	r := csv.NewReader(reader)
	r.Comma = '|'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	var header []string
	firstLine := true
	for {
		if firstLine == false {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if header == nil {
				header = record
			} else {
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
					EstadoISO:    rows[isoCode]["isoCode"],
					Municipio:    dict["D_mnpio"],
					Ciudad:       dict["d_ciudad"],
					Barrio:       dict["d_asenta"],
				}
				zipCodes = append(zipCodes, parsingZipCode)
				zipCodesModel = append(zipCodesModel, parsingZipCode)
			}
		} else {
			r.Read()
			firstLine = false
		}
	}
	return zipCodes, zipCodesModel
}
