package main

import (
	"github.com/Prashant-Mediastinct/Assignment/adservice/controllers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/adunit/", controllers.GetPublisherId)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
