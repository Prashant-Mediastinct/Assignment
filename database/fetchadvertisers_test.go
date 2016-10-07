package database

import "testing"
import "github.com/Prashant-Mediastinct/Assignment/models"

func TestFetchAdvertiser(t *testing.T) {

	var Advertiser []models.AdvertiserData
	params := "10"

	expected := "11"

	DBDef()

	Advertiser = FetchAdvertisers(params)

	if Advertiser[0].Advertiser_id != expected {
		t.Errorf("Publisher_id : got %s want %s", Advertiser[0].Advertiser_id, expected)
	}

}
