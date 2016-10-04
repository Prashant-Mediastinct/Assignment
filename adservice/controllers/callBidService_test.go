package controllers

import (
	//"github.com/Prashant-Mediastinct/Assignment/adservice/controllers"
	"log"
	"testing"
)

func TestBidService(t *testing.T) {

	log.Println("Testing for BidService")
	var Publisher_Id = "4"
	var params = "4"

	adunit := callBidService(Publisher_Id, params)

	log.Printf(" Data : %+v", adunit)

	log.Println("BidService testing : SUCCESSFUL")
}
