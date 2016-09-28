package main

import (
	"github.com/Prashant-Mediastinct/Assignment/bidservice/controllers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/publisher_id/", controllers.GetAdvertisers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
