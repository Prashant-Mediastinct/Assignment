package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Prashant-Mediastinct/Assignment/models"
	as "github.com/aerospike/aerospike-client-go"
	"log"
	"net/http"
	//"reflect"
)

func checkAerospikeForPublisherId(params string, w http.ResponseWriter) bool {

	var adunit models.AdunitData
	keyName := fmt.Sprintf("adUnit:%s", params)
	Key1, _ = as.NewKey("test", "myset", keyName)

	Adunit := &adunit

	err := Client.GetObject(ReadPolicy, Key1, Adunit)

	//Adunit := record
	//log.Println(reflect.TypeOf(adunit))

	if err == nil {

		log.Println("Fetched from Aerospike!!")
		log.Printf("Record found : %+v", Adunit)

		log.Printf("Adunit: %+v ", Adunit)
		json.NewEncoder(w).Encode(Adunit)
		return true
	} else {
		return false
	}

	//return false
}
