package cache

import (
	"encoding/json"
	"fmt"
	"github.com/Prashant-Mediastinct/Assignment/models"
	as "github.com/aerospike/aerospike-client-go"
	"log"
	"net/http"
)

var (
	Key         *as.Key
	WritePolicy *as.WritePolicy
	ReadPolicy  *as.BasePolicy
)

type CacheClient struct {
	client *as.Client
}

/*type ClientConfig {
	Host string
	Port int
}*/

func init() {

	var cc *CacheClient
	var err error
	cc.client, err = as.NewClient("127.0.0.1", 3000)

	if !cc.client.IsConnected() {
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

func (cc *CacheClient) GetDataFromCache(params string, w http.ResponseWriter) bool {

	var adunit models.AdunitData
	keyName := fmt.Sprintf("adUnit:%s", params)
	Key, _ = as.NewKey("test", "myset", keyName)

	Adunit := &adunit

	err := cc.client.GetObject(ReadPolicy, Key, Adunit)

	if err == nil {
		log.Println("Fetched from Aerospike!!")
		log.Printf("Record found : %+v", Adunit)

		log.Printf("Adunit: %+v ", Adunit)
		json.NewEncoder(w).Encode(Adunit)
		return true
	} else {
		log.Println("No Such Data Exists!")
		return false
	}
}

func (cc *CacheClient) PostDataIntoCache(Adunit models.AdunitData) {

	adunit := &Adunit
	err := cc.client.PutObject(WritePolicy, Key, adunit)

	if err != nil {
		log.Fatalf("Error  :", err.Error())
	}
}
