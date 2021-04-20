package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MYSQLConnection bla bla...
func MYSQLConnection() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root1"
	dbPass := "root1"
	dbName := "teachers_hcn"
	dbURL := "192.168.120.42:3306"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbURL+")/"+dbName)
	return db, err
}
