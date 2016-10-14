package database

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func FetchPublisher(adUnit_Id string) string {

	var publisher_Id string
	row := DB.QueryRow("select Publisher_id from Adunits where Adunit_id=?", adUnit_Id)

	err := row.Scan(&publisher_Id)
	if err != nil {
		log.Println("No Data for Adunit_Id: ", adUnit_Id)
	}

	return publisher_Id
}
