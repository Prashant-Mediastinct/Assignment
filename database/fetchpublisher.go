package database

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func FetchPublisher(params string) string {

	var Publisher_Id string
	row := DB.QueryRow("select Publisher_id from Adunits where Adunit_id=?", params)

	log.Println("%+v", row)
	err := row.Scan(&Publisher_Id)
	if err != nil {
		log.Fatal(err)
	}

	return Publisher_Id
}
