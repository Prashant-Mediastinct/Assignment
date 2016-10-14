package controllers

import (
	"encoding/json"
	"github.com/Prashant-Mediastinct/Assignment/models"
	"log"
	"net/http"
)

func callBidService(Publisher_id, params string) models.AdunitData {

	var Advertiser []models.AdvertiserData

	var Adunit models.AdunitData
	url := "http://localhost:8081/publisher_id/" + Publisher_id

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(res.Body).Decode(&Advertiser)
	//log.Println("Response body : ", res.Body)

	//log.Println("Data: ", Advertiser)

	Adunit.Adunit_id = params
	Adunit.Adunit_info.Advertiser = Advertiser

	return Adunit
}
