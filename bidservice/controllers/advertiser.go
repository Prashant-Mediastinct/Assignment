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

	publisher_Id := req.URL.Path[14:]
	log.Println("Publisher_id passed : ", publisher_Id)

	Advertiser = database.FetchAdvertisers(publisher_Id)

	json.NewEncoder(w).Encode(Advertiser)
}
