package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Prashant-Mediastinct/Assignment/database"
	as "github.com/aerospike/aerospike-client-go"
	"log"
	"net/http"
)

var Publisher_Id string
var Client *as.Client
var Key, Key1 *as.Key
var WritePolicy *as.WritePolicy
var ReadPolicy *as.BasePolicy

//var Adunit models.AdunitData

func init() {

	var err error
	Client, err = as.NewClient("127.0.0.1", 3000)

	if !Client.IsConnected() {
		log.Fatal("Not Connected!")
	} else {
		log.Println("Connected!!")
	}

	if err != nil {
		log.Fatalf("Error  :", err.Error())
	}

	WritePolicy = as.NewWritePolicy(0, 10)
	ReadPolicy = as.NewPolicy()

}

func GetPublisherId(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Path[8:9]
	log.Println("Params passed : ", params)

	keyName := fmt.Sprintf("adUnit:%s", params)
	Key, _ = as.NewKey("test", "myset", keyName)

	status := checkAerospikeForPublisherId(params, w)

	if status != true {
		database.DBDef()

		log.Println("Fetching From database!!")
		Publisher_Id = database.FetchPublisher(params)

		Adunit := callBidService(Publisher_Id, params)

		adunit := &Adunit
		err := Client.PutObject(WritePolicy, Key, adunit)

		if err != nil {
			log.Fatalf("Error  :", err.Error())
		}
		json.NewEncoder(w).Encode(Adunit)
	}
}
