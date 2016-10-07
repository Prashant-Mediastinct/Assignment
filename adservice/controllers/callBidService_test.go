package controllers

import "testing"
import "github.com/Prashant-Mediastinct/Assignment/models"

func TestBidService(t *testing.T) {

	var adunit models.AdunitData
	var Publisher_Id = "10"
	var params = "7"
	var expected = "7"

	adunit = callBidService(Publisher_Id, params)

	if adunit.Adunit_id != expected {
		t.Errorf("Adunit_id : got %s want %s", adunit.Adunit_id, expected)
	}

}
