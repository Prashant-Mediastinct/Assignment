package main

import (
	"github.com/Prashant-Mediastinct/Assignment/adservice/controllers"
	"github.com/Prashant-Mediastinct/Assignment/database"
	"log"
	"net/http"
)

func main() {

	database.DBDef()
	defer database.DB.Close()

	http.HandleFunc("/adunit/", controllers.GetPublisherId)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
