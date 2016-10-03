package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func initDB(dataSourceName string) *sql.DB {

	DB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal()
	}

	//check for availability..
	err1 := DB.Ping()
	if err1 != nil {
		log.Fatal()
	}
	return DB
}
