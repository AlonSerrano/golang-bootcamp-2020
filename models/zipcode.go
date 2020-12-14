package models

// ZipCodeCSV this structure is the base to define the zipcodes
type ZipCodeCSV struct {
	Id           string `bson:"_id"`
	CodigoPostal string `bson:"codigoPostal"`
	Estado       string `bson:"estado"`
	EstadoISO    string `bson:"estadoISO"`
	Municipio    string `bson:"municipio"`
	Ciudad       string `bson:"ciudad,omitempty"`
	Barrio       string `bson:"barrio"`
}
