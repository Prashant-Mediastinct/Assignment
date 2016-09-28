package controllers

import (
	"encoding/json"
	"github.com/Prashant-Mediastinct/Assignment/databaseconn"
	"github.com/Prashant-Mediastinct/Assignment/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var Advertise []models.AdvertiserData

func GetAdvertisers(w http.ResponseWriter, req *http.Request) {

	var ad models.AdvertiserData
	var cda []models.CreativeData

	db := databaseconn.InitDB("root:root@/test")

	defer db.Close()

	//Extract the Publisher_id and fetch the corresponding rows from Advertisers table...
	params := req.URL.Path[14:]
	log.Println("Publisher_id passed : ", params)
	rows_advertiser, err := db.Query("select Advertiser_id, Advertise_name from Advertisers where Publisher_id=? and status = 'Active'", params)

	for rows_advertiser.Next() {
		err = rows_advertiser.Scan(&ad.Advertiser_id, &ad.Advertise_name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Advertiser_id: ", ad.Advertiser_id)
		log.Println("Advertiser_name: ", ad.Advertise_name)

		cda = GetCreatives(ad.Advertiser_id)

		ad.CreativeData = &cda
		Advertise = append(Advertise, ad)
	}
	log.Printf("%+v", Advertise[0].CreativeData)
	json.NewEncoder(w).Encode(Advertise)
}

func GetCreatives(Advertiser_id int) []models.CreativeData {

	var cda []models.CreativeData
	var cd models.CreativeData

	db := databaseconn.InitDB("root:root@/test")

	defer db.Close()

	rows_creatives, err := db.Query("select Creative_id, Content from Creatives where Advertiser_id=?", Advertiser_id)
	for rows_creatives.Next() {
		err = rows_creatives.Scan(&cd.Creative_id, &cd.Content)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Creatives_id: ", cd.Creative_id)
		log.Println("Content: ", cd.Content)
		cda = append(cda, cd)
		log.Println(cda)
	}
	return cda
}
