package databaseconn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDB(dataSourceName string) *sql.DB {

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	//check for availability..
	err1 := db.Ping()
	if err1 != nil {
		log.Fatal(err1)
	}
	return db

}
