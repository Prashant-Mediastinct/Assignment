package database

import (
	"github.com/Prashant-Mediastinct/Assignment/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func fetchCreatives(Advertiser_id string) []models.CreativeData {

	var cda []models.CreativeData
	var cd models.CreativeData

	//log.Printf("DB : %+v", db)
	//log.Println(Advertiser_id)
	rows_creatives, err := DB.Query("select Creative_id, Content from Creatives where Advertiser_id=?", Advertiser_id)
	//log.Println("Here!")
	//log.Println(Advertiser_id)

	//log.Printf("here : %+v", rows_creatives)

	for rows_creatives.Next() {
		err = rows_creatives.Scan(&cd.Creative_id, &cd.Content)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println("Creatives_id: ", cd.Creative_id)
		//log.Println("Content: ", cd.Content)
		cda = append(cda, cd)
		//log.Println(cda)
	}
	return cda
}
