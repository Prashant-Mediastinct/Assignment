package database

import (
	//"github.com/Prashant-Mediastinct/Assignment/database"
	"log"
	"testing"
)

func TestFetchAdvertiser(t *testing.T) {

	params := "4"

	DBDef()

	Advertiser := FetchAdvertisers(params)

	log.Println("Advertiser Data : ", Advertiser)
}
