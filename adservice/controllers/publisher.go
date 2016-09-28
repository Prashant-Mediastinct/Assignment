package controllers

import (
	"github.com/Prashant-Mediastinct/Assignment/databaseconn"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var Publisher_id string

func GetPublisherId(w http.ResponseWriter, req *http.Request) {

	db := databaseconn.InitDB("root:root@/test")

	defer db.Close()

	params := req.URL.Path[8:9]
	log.Println(params)
	row := db.QueryRow("select Publisher_id from Adunits where Adunit_id=?", params)

	err := row.Scan(&Publisher_id)
	log.Println(Publisher_id)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Publisher ID: ", Publisher_id)

	url := "http://localhost:8081/publisher_id/" + Publisher_id

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
