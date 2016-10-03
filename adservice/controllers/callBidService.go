package controllers

import (
	"encoding/json"
	"github.com/Prashant-Mediastinct/Assignment/models"
	"log"
	"net/http"
)

func callBidService(Publisher_id, params string) {

	var Advertiser []models.AdvertiserData

	var Adunit models.AdunitData
	url := "http://localhost:8081/publisher_id/" + Publisher_id

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(res.Body).Decode(&Advertiser)
	//log.Println("Response body : ", res.Body)

	log.Println("Data: ", Advertiser)

	AdunitData.Adunit_id = params
	/*log.Printf("Hello there : %+v", res)

	if res == nil {
		log.Println("No advertiser available for the given publisher id")
	} else {
		log.Printf("Response body : ", res)
	}*/
}
