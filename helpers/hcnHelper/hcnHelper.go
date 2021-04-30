package hcnHelper

import (
	"hcn/config"
	"hcn/mymodels"
)

// CreateHCN creates an HCN in MySQl db
func CreateHCN(hcn mymodels.HCN) (mymodels.HCN, error) {
	var emptyHCN mymodels.HCN
	var Db, err = config.MYSQLConnection()
	defer Db.Close()
	if err != nil {
		return emptyHCN, err
	}
	rows, err := Db.Exec("INSERT INTO HCN(TeacherID, MongoID) VALUES (?, ?)", hcn.TeacherID, hcn.MongoID)
	if err != nil {
		return emptyHCN, err
	}
	cnt, err := rows.RowsAffected()
	if cnt == 1 {
		int64ID, _ := rows.LastInsertId()
		intID := int(int64ID)
		hcn.ID = &intID
		return hcn, nil
	} else {
		return emptyHCN, err
	}
}

// GetSqlID returns Mongo ID from SQL ID
func GetSqlID(sqlID int) (string, error) {
	var mongoID string
	var Db, _ = config.MYSQLConnection()
	err := Db.QueryRow("SELECT MongoID FROM HCN WHERE ID=?", sqlID).Scan(&mongoID)
	defer Db.Close()
	if err != nil {
		return mongoID, err
	}
	return mongoID, err
}
