package main

import (
	as "github.com/aerospike/aerospike-client-go"
	"log"
)

func AerospikeAddData() {

	client, err := as.NewClient("127.0.0.1", 3000)

	if !client.IsConnected() {
		log.Fatal("Not Connected!")
	} else {
		log.Println("Connected!!")
	}

	if err != nil {
		log.Fatalf("Error  :", err.Error())
	}

	key, _ := as.NewKey("test", "myset", "mykey")
	bin1 := as.NewBin("adunit", "7")
	bin2 := as.NewBin("publisher_id", "10")

	client.PutBins(nil, key, bin1, bin2)

	record, err := client.Get(nil, key, "adunit", "publisher_id")

	log.Printf("Record : %+v", record)
}
