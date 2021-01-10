package myhandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateClinicalCase bla bla...
func CreateClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClinicalCase ClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newClinicalCase)
	switch {
	case (*newClinicalCase.TeachersID*1 == 0) || (*newClinicalCase.TeachersID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid TeachersID", *newClinicalCase.TeachersID)
		return
	case (newClinicalCase.Title == nil) || (len(*newClinicalCase.Title) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case (newClinicalCase.Description == nil) || (len(*newClinicalCase.Description) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case (newClinicalCase.Media == nil) || (len(*newClinicalCase.Media) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Media is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Clinical_Cases(Title,Description,Media,TeachersId) VALUES (?,?,?,?)", newClinicalCase.Title, newClinicalCase.Description, newClinicalCase.Media, newClinicalCase.TeachersID)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, err := rows.RowsAffected()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		} else if cnt < 1 {
			fmt.Fprintf(w, "No rows affected")
			return
		}
		lastID, err := rows.LastInsertId()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		fmt.Fprintf(w, "ID inserted: %v", lastID)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// GetClinicalCases bla bla...
func GetClinicalCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clinicalCases AllClinicalCases
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT Id, Title, Description, Media, TeachersId FROM Clinical_Cases")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, TeachersID int
		var Title, Description, Media string
		if err := rows.Scan(&ID, &Title, &Description, &Media, &TeachersID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var clinicalCase = ClinicalCase{ID: &ID, Title: &Title, Description: &Description, Media: &Media, TeachersID: &TeachersID}
		clinicalCases = append(clinicalCases, clinicalCase)
	}
	json.NewEncoder(w).Encode(clinicalCases)
	w.WriteHeader(http.StatusOK)
	return
}

// GetClinicalCase bla bla...
func GetClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	clinicalCaseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var ID, TeachersID int
	var Title, Description, Media string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID,Title,Description,Media,TeachersId FROM Clinical_Cases WHERE Id=?", clinicalCaseID).Scan(&ID, &Title, &Description, &Media, &TeachersID)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var clinicalCase = ClinicalCase{ID: &ID, Title: &Title, Description: &Description, Media: &Media, TeachersID: &TeachersID}
	json.NewEncoder(w).Encode(clinicalCase)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateClinicalCase bla bla...
func UpdateClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedClinicalCase ClinicalCase
	json.Unmarshal(reqBody, &updatedClinicalCase)
	switch {
	case (updatedClinicalCase.ID) == nil || (*updatedClinicalCase.ID*1 == 0) || (*updatedClinicalCase.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case updatedClinicalCase.Title == nil || len(*updatedClinicalCase.Title) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case updatedClinicalCase.Description == nil || len(*updatedClinicalCase.Description) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case updatedClinicalCase.Media == nil || len(*updatedClinicalCase.Media) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Media is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Clinical_Cases SET Title=?, Description=?, Media=? WHERE Id=?", updatedClinicalCase.Title, updatedClinicalCase.Description, updatedClinicalCase.Media, updatedClinicalCase.ID)
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
			fmt.Fprintf(w, "One row updated")
		} else {
			fmt.Fprintf(w, "No rows updated")
		}

		w.Header().Set("Content-Type", "application/json")
		return
	}
}

// DeleteClinicalCase bla bla...
func DeleteClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedClinicalCase ClinicalCase
	json.Unmarshal(reqBody, &deletedClinicalCase)

	if (deletedClinicalCase.ID) == nil || (*deletedClinicalCase.ID*1 == 0) || (*deletedClinicalCase.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Clinical_Cases WHERE Id=?", deletedClinicalCase.ID)
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
