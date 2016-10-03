package controllers

import (
	"log"
	"net/http"
)

func callBidService(Publisher_id string) {

	url := "http://localhost:8081/publisher_id/" + Publisher_id

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
