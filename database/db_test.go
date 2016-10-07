package database

import "testing"
import "database/sql"

func TestDBConn(t *testing.T) {

	var DB *sql.DB
	DB = initDB("root:root@/test")

	err1 := DB.Ping()
	if err1 != nil {
		t.Error("Database Unavailable")
	}
}
