package database

import (
	//"github.com/Prashant-Mediastinct/Assignment/database"
	"log"
	"testing"
)

func TestFetchPublishers(t *testing.T) {

	params := "5"

	DBDef()

	Publisher_Id := FetchPublisher(params)

	if Publisher_Id == "" {
		t.Errorf("%s Doesnt exist!", Publisher_Id)
	} else {
		log.Println("Advertiser Data : ", Publisher_Id)
	}
}
