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

	params := req.URL.Path[14:]
	log.Println("Publisher_id passed : ", params)

	Advertiser = database.FetchAdvertisers(params)

	json.NewEncoder(w).Encode(Advertiser)
}
