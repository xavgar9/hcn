package hcn

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strconv"
)

// CreateHCN bla bla...
func CreateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newHCN mymodels.HCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newHCN)
	switch {
	case (newHCN.TeacherID == nil) || (*newHCN.TeacherID*1 == 0) || (*newHCN.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO HCN(TeacherID) VALUES (?)", newHCN.TeacherID)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 1 {
			int64ID, _ := rows.LastInsertId()
			intID := int(int64ID)
			newHCN.ID = &intID
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newHCN)
		}
		return
	}
}

// GetAllHCN bla bla...
func GetAllHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hcns mymodels.AllHCN
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, TeacherID FROM HCN")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, TeacherID int
		if err := rows.Scan(&ID, &TeacherID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var hcn = mymodels.HCN{ID: &ID, TeacherID: &TeacherID}
		hcns = append(hcns, hcn)
	}
	json.NewEncoder(w).Encode(hcns)
	w.WriteHeader(http.StatusOK)
	return
}

// GetHCN bla bla...
func GetHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
		return
	}
	hcnID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
		return
	}
	var ID, TeacherID int
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, TeacherID FROM HCN WHERE ID=?", hcnID).Scan(&ID, &TeacherID)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var course = mymodels.HCN{ID: &ID, TeacherID: &TeacherID}
	json.NewEncoder(w).Encode(course)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateHCN bla bla...
func UpdateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedHCN mymodels.HCN
	json.Unmarshal(reqBody, &updatedHCN)
	switch {
	case (updatedHCN.ID == nil) || (*updatedHCN.ID*1 == 0) || (*updatedHCN.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedHCN.TeacherID == nil) || (*updatedHCN.TeacherID*1 == 0) || (*updatedHCN.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE HCN SET TeacherID=? WHERE ID=?", updatedHCN.TeacherID, updatedHCN.ID)
		defer Db.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}

		count, err := row.RowsAffected()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		if count == 1 {
			json.NewEncoder(w).Encode(updatedHCN)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}

// DeleteHCN bla bla...
func DeleteHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedHCN mymodels.HCN
	json.Unmarshal(reqBody, &deletedHCN)

	if (deletedHCN.ID) == nil || (*deletedHCN.ID*1 == 0) || (*deletedHCN.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM HCN WHERE ID=?", deletedHCN.ID)
	defer Db.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		return
	}

	count, err := row.RowsAffected()
	if err != nil {
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		return
	}
	if count == 1 {
		fmt.Fprintf(w, "One row deleted")
	} else {
		fmt.Fprintf(w, "No rows deleted")
	}
	w.Header().Set("Content-Type", "application/json")
	return
}
