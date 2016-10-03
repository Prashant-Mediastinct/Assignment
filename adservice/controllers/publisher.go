package controllers

import (
	//  "github.com/Prashant-Mediastinct/Assignment/bidservice/controllers"
	"github.com/Prashant-Mediastinct/Assignment/database"
	//"log"
	"net/http"
)

var Publisher_Id string

func GetPublisherId(w http.ResponseWriter, req *http.Request) {

	database.DBDef()

	params := req.URL.Path[8:9]
	//log.Println(params)

	Publisher_Id = database.FetchPublisher(params)
	//log.Println(Publisher_Id)

	//log.Println("Publisher ID: ", Publisher_Id)

	callBidService(Publisher_Id, params)
}
