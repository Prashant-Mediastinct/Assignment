package main

import (
	"log"
	"net/http"

	"github.com/Prashant-Mediastinct/Assignment/bidservice/controllers"
	"github.com/Prashant-Mediastinct/Assignment/database"
)

func main() {

	database.DBDef()
	defer database.DB.Close()

	http.HandleFunc("/publisher_id/", controllers.GetAdvertisers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
