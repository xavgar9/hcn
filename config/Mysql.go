package config

import (
	"database/sql"
)

// MYSQLConnection bla bla...
func MYSQLConnection() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "teachers_hcn"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
