package database

import (
	"github.com/Prashant-Mediastinct/Assignment/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func FetchAdvertisers(params string) []models.AdvertiserData {

	var Advertiser []models.AdvertiserData

	var ad models.AdvertiserData
	var cda []models.CreativeData
	rows_advertiser, err := DB.Query("select Advertiser_id, Advertise_name from Advertisers where Publisher_id=? and status = 'Active'", params)

	for rows_advertiser.Next() {
		err = rows_advertiser.Scan(&ad.Advertiser_id, &ad.Advertiser_name)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println("Advertiser_id: ", ad.Advertiser_id)
		//log.Println("Advertiser_name: ", ad.Advertiser_name)

		cda = fetchCreatives(ad.Advertiser_id)

		ad.CreativeData = cda
		Advertiser = append(Advertiser, ad)
	}
	//log.Printf("%+v", Advertiser[0].CreativeData)

	return Advertiser
}
