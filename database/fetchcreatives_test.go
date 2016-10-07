package database

import "testing"
import "github.com/Prashant-Mediastinct/Assignment/models"

func TestFetchCreatives(t *testing.T) {

	var CDA []models.CreativeData
	params := "11"
	expected := "11"
	DBDef()

	CDA = fetchCreatives(params)

	if CDA[0].Creative_id != expected {
		t.Errorf("Publisher_id : got %s want %s", CDA[0].Creative_id, expected)
	}
	//log.Println("Advertiser Data : ", CDA)
}
