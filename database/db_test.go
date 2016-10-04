package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestDBConn(t *testing.T) {

	dataSourceName := "root:root@/test"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
		log.Println(err.Error())
	}
	log.Println("DB connection opened...")

	//check for availability..
	err1 := db.Ping()
	if err1 != nil {
		log.Fatal(err1)
		log.Println(err1.Error())
	}
	log.Println("DB available...")
}
