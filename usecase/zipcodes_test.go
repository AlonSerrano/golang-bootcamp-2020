package usecase

import (
	"fmt"
	"testing"

	"github.com/AlonSerrano/GolangBootcamp/service"

	"github.com/stretchr/testify/assert"
)

func Test_getCSVCodes(t *testing.T) {
	//TODO mock of mongo
	collection := service.UseZipCodeTable()
	assert.NotEqual(t, 0, len(GetAndSave(collection).InsertedIDs), "Records have been found")
}

func Test_searchZipCodes(t *testing.T) {
	//TODO mock of mongo
	collection := service.UseZipCodeTable()
	assert.Condition(t, func() bool {
		r := SearchZipCodes("97306", collection)
		fmt.Println(r)
		if len(r) != 0 {
			return true
		}
		return false
	}, "Result of search > 0")
}

func Test_dropZipCodes(t *testing.T) {
	collection := service.UseZipCodeTable()
	assert.Equal(t, true, dropZipCodes(collection), "")
}
