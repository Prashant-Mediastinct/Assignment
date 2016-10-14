package database

import (
	"github.com/Prashant-Mediastinct/Assignment/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func FetchAdvertisers(publisher_Id string) []models.AdvertiserData {

	var Advertiser []models.AdvertiserData

	var ad models.AdvertiserData
	var cda []models.CreativeData

	rows_advertiser, err := DB.Query("select Advertiser_id, Advertise_name from Advertisers where Publisher_id=? and status = 'Active'", publisher_Id)

	for rows_advertiser.Next() {
		err = rows_advertiser.Scan(&ad.Advertiser_id, &ad.Advertiser_name)
		if err != nil {
			log.Println("No data")
			log.Fatal(err)
		}

		cda = fetchCreatives(ad.Advertiser_id)

		ad.CreativeData = cda
		Advertiser = append(Advertiser, ad)
	}

	return Advertiser
}
