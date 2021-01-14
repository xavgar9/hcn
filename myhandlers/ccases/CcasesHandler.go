package ccases

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/helpers"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateClinicalCase bla bla...
func CreateClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClinicalCase mymodels.ClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newClinicalCase)
	switch {
	case (*newClinicalCase.TeacherID*1 == 0) || (*newClinicalCase.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid TeacherID", *newClinicalCase.TeacherID)
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
		rows, err := Db.Exec("INSERT INTO Clinical_Cases(Title,Description,Media,TeacherID) VALUES (?,?,?,?)", newClinicalCase.Title, newClinicalCase.Description, newClinicalCase.Media, newClinicalCase.TeacherID)
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

// GetAllClinicalCases bla bla...
func GetAllClinicalCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clinicalCases mymodels.AllClinicalCases
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, Title, Description, Media, TeacherID FROM Clinical_Cases")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, TeacherID int
		var Title, Description, Media string
		if err := rows.Scan(&ID, &Title, &Description, &Media, &TeacherID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var clinicalCase = mymodels.ClinicalCase{ID: &ID, Title: &Title, Description: &Description, Media: &Media, TeacherID: &TeacherID}
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
	var ID, TeacherID int
	var Title, Description, Media string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID,Title,Description,Media,TeacherID FROM Clinical_Cases WHERE ID=?", clinicalCaseID).Scan(&ID, &Title, &Description, &Media, &TeacherID)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var clinicalCase = mymodels.ClinicalCase{ID: &ID, Title: &Title, Description: &Description, Media: &Media, TeacherID: &TeacherID}
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
	var updatedClinicalCase mymodels.ClinicalCase
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
		row, err := Db.Exec("UPDATE Clinical_Cases SET Title=?, Description=?, Media=? WHERE ID=?", updatedClinicalCase.Title, updatedClinicalCase.Description, updatedClinicalCase.Media, updatedClinicalCase.ID)
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
	var deletedClinicalCase mymodels.ClinicalCase
	json.Unmarshal(reqBody, &deletedClinicalCase)

	if (deletedClinicalCase.ID) == nil || (*deletedClinicalCase.ID*1 == 0) || (*deletedClinicalCase.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Clinical_Cases WHERE ID=?", deletedClinicalCase.ID)
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

// AddHCN adds an HCN into a Clinical Case...
func AddHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var newHCNVinculation mymodels.HCNVinculation
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newHCNVinculation)
	switch {
	case (newHCNVinculation.ClinicalCaseID == nil) || (*newHCNVinculation.ClinicalCaseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (newHCNVinculation.HCNID == nil) || (*newHCNVinculation.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO CCases_HCN(ClinicalCaseID,HCNID) VALUES (?,?)", newHCNVinculation.ClinicalCaseID, newHCNVinculation.HCNID)
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

// RemoveHCN from a Clnical Case...
func RemoveHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var removeHCNVinculation mymodels.HCNVinculation
	json.Unmarshal(reqBody, &removeHCNVinculation)
	if (removeHCNVinculation.ClinicalCaseID) == nil || (*removeHCNVinculation.ClinicalCaseID*1 == 0) || (*removeHCNVinculation.ClinicalCaseID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	}
	if (removeHCNVinculation.HCNID) == nil || (*removeHCNVinculation.HCNID*1 == 0) || (*removeHCNVinculation.HCNID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM CCases_HCN WHERE ClinicalCaseID=? AND HCNID=?", removeHCNVinculation.ClinicalCaseID, removeHCNVinculation.HCNID)
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
