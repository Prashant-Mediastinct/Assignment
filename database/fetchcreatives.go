package database

import (
	"github.com/Prashant-Mediastinct/Assignment/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func fetchCreatives(advertiser_Id string) []models.CreativeData {

	var cda []models.CreativeData
	var cd models.CreativeData

	rows_creatives, err := DB.Query("select Creative_id, Content from Creatives where Advertiser_id=?", advertiser_Id)

	for rows_creatives.Next() {
		err = rows_creatives.Scan(&cd.Creative_id, &cd.Content)
		if err != nil {
			log.Fatal(err)
		}

		cda = append(cda, cd)
	}
	return cda
}
