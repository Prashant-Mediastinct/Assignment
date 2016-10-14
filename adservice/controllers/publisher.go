package controllers

import (
	"encoding/json"
	"github.com/Prashant-Mediastinct/Assignment/cache"
	"github.com/Prashant-Mediastinct/Assignment/database"
	"github.com/Prashant-Mediastinct/Assignment/models"
	"log"
	"net/http"
	"strings"
)

var cc *cache.CacheClient

func init() {
	cc = cc.InitialiseCacheClient()
}

func GetPublisherId(w http.ResponseWriter, req *http.Request) {

	var dataChan chan models.AdunitData
	var signalChan chan struct{}
	var Adunit models.AdunitData

	params := req.URL.Path[8:]

	paramsArray := strings.Split(params, ",")
	log.Println("Query passed : ", paramsArray)
	newParamsArray := removeDuplicates(paramsArray)
	log.Println("Optimised Query : ", newParamsArray)

	// NOTE : instance of dataChan should be created in the caller func. So that only one instance is used throughout.
	dataChan = make(chan models.AdunitData)
	signalChan = make(chan struct{})

	for i := 0; i < len(newParamsArray); i++ {
		go getData(newParamsArray[i], Adunit, dataChan, signalChan)
	}

	Adunit = <-dataChan
	json.NewEncoder(w).Encode(Adunit)

	log.Println("AdUnit received from datachan")
}

func getData(adUnit_Id string, Adunit models.AdunitData, dataChan chan models.AdunitData, signalChan chan struct{}) {

	//defer wg.Done()
	var status bool
	var publisher_Id string
	//log.Println("Here!")
	Adunit, status = cc.GetDataFromCache(adUnit_Id)

	if status != true {
		database.DBDef()

		//log.Println("Fetching From database for AdUnit_Id : ", adUnit_Id)
		publisher_Id = database.FetchPublisher(adUnit_Id)
		if publisher_Id == "" {
			//time.Sleep(100 * time.Millisecond)
			return
		}
		//log.Printf("Publisher_Id : %s", Publisher_Id)
		Adunit = callBidService(publisher_Id, adUnit_Id)
		//log.Printf("Adunit : %#v", Adunit)
		cc.PostDataIntoCache(Adunit)
	}

	select {
	case dataChan <- Adunit:
		close(signalChan)
		log.Printf("AdUnit fetch for Adunit_id : %s", Adunit.Adunit_id)
		log.Printf("Data: %+v", Adunit)

	case <-signalChan:
		log.Printf("Terminating AdUnit fetch for Adunit_id : %s", Adunit.Adunit_id)
		return
	}
}

func removeDuplicates(paramsArray []string) []string {

	encountered := map[string]bool{}
	result := []string{}

	for n := range paramsArray {

		if encountered[paramsArray[n]] == true {
			//ignore
		} else {
			encountered[paramsArray[n]] = true
			result = append(result, paramsArray[n])
		}
	}
	return result
}
