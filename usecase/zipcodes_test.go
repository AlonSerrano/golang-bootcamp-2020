package usecase

import (
	"fmt"
	"testing"

	"github.com/AlonSerrano/GolangBootcamp/service"

	"github.com/stretchr/testify/assert"
)

func Test_getCSVCodes(t *testing.T) {
	//TODO mock de mongo
	client := service.GetConnDB()
	assert.NotEqual(t, 0, len(GetAndSave(client).InsertedIDs), "Records have been found")
}

func Test_searchZipCodes(t *testing.T) {
	client := service.GetConnDB()
	assert.Condition(t, func() bool {
		r := SearchZipCodes("97306", client)
		fmt.Println(r)
		if len(r) != 0 {
			return true
		}
		return false
	}, "Result of search > 0")
}

func Test_dropZipCodes(t *testing.T) {
	client := service.GetConnDB()
	assert.Equal(t, true, dropZipCodes(client), "")
}
