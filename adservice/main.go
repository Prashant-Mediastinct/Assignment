package main

import (
	"github.com/Prashant-Mediastinct/Assignment/adservice/controllers"
	"github.com/Prashant-Mediastinct/Assignment/database"
	"log"
	"net/http"
	//"reflect"
)

/*var Client *as.Client
var Key *as.Key

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
}*/

func main() {

	/*Key, _ = as.NewKey("test", "myset", "mykey")
	bin1 := as.NewBin("adunit", "7")
	bin2 := as.NewBin("publisher_id", "10")

	Client.PutBins(nil, Key, bin1, bin2)

	record, err := Client.Get(nil, Key, "adunit", "publisher_id")

	if err != nil {
		log.Fatalf("Error  :", err.Error())
	}

	log.Printf("Record : %+v", record.Bins["adunit"])
	*/
	database.DBDef()
	defer database.DB.Close()

	http.HandleFunc("/adunit/", controllers.GetPublisherId)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//Client.Close()
}
