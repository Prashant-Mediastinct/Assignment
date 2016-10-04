package database

import (
	//"github.com/Prashant-Mediastinct/Assignment/database"
	"log"
	"testing"
)

func TestFetchCreatives(t *testing.T) {

	params := 4

	DBDef()

	CDA := fetchCreatives(params)

	log.Println("Advertiser Data : ", CDA)
}
