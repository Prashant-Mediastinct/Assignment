package controllers

import (
	"encoding/json"
	"github.com/Prashant-Mediastinct/Assignment/database"
	"github.com/Prashant-Mediastinct/Assignment/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func GetAdvertisers(w http.ResponseWriter, req *http.Request) {

	var Advertiser []models.AdvertiserData
	//_ = Advertiser
	params := req.URL.Path[14:]
	log.Println("Publisher_id passed : ", params)

	if params < "7" {
		Advertiser = database.FetchAdvertisers(params)
	} else {
		log.Printf("Record with Publisher_Id : %s doesn't exist", params)
	}
	json.NewEncoder(w).Encode(Advertiser)
}
