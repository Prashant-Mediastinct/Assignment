package database

import (
	"database/sql"
)

var DB *sql.DB

func DBDef() {

	DB = initDB("root:root@/test")
}
