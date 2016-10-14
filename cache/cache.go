package cache

import (
	"fmt"
	"github.com/Prashant-Mediastinct/Assignment/models"
	as "github.com/aerospike/aerospike-client-go"
	"log"
)

var (
	Key         *as.Key
	WritePolicy *as.WritePolicy
	ReadPolicy  *as.BasePolicy
)

type CacheClient struct {
	client *as.Client
}

func (cc *CacheClient) InitialiseCacheClient() *CacheClient {

	var err error

	//log.Printf("CC : %#v", CC)
	cc = &CacheClient{}
	cc.client, err = as.NewClient("127.0.0.1", 3000)

	//log.Printf("ClientCache : %+v", CC)

	if err != nil {
		log.Fatalf("Error  :", err.Error())
	}

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
	return cc
}

func (cc *CacheClient) GetDataFromCache(adUnit string) (models.AdunitData, bool) {

	var adunit models.AdunitData
	keyName := fmt.Sprintf("adUnit:%s", adUnit)
	Key, _ = as.NewKey("test", "myset", keyName)
	//log.Println("In getdatafromcache")
	err := cc.client.GetObject(ReadPolicy, Key, &adunit)

	if err == nil {
		log.Println("Fetched from Aerospike for AdUnit_Id : ", adUnit)
		return adunit, true
	} else {
		return adunit, false
	}
}

func (cc *CacheClient) PostDataIntoCache(Adunit models.AdunitData) {

	//adunit := &Adunit
	err := cc.client.PutObject(WritePolicy, Key, &Adunit)

	if err != nil {
		log.Fatalf("Error  :", err.Error())
	}
}
