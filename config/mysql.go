package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MYSQLConnection bla bla...
func MYSQLConnection() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "teachers_hcn"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
