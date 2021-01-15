package courses

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

// AddClinicalCase adds a relationship between a course and a HCN...
func AddClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourseClinicalCase mymodels.CourseClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newCourseClinicalCase)
	switch {
	case (*newCourseClinicalCase.CourseID*1 == 0) || (*newCourseClinicalCase.CourseID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (*newCourseClinicalCase.ClinicalCaseID*1 == 0) || (*newCourseClinicalCase.ClinicalCaseID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (*newCourseClinicalCase.Displayable*1 == 0) || (*newCourseClinicalCase.Displayable*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid display status", *newCourseClinicalCase.Displayable)
		return
	default:
		rows, err := Db.Exec("INSERT INTO Courses_HCN(CourseID,ClinicalCaseID,Displayable) VALUES (?,?,?)", newCourseClinicalCase.CourseID, newCourseClinicalCase.ClinicalCaseID, newCourseClinicalCase.Displayable)
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
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}

	var allCourseClinicalCase mymodels.AllCourseClinicalCase
	var Db, _ = config.MYSQLConnection()
	//rows, err := Db.Query("SELECT ID, TeacherID FROM HCN WHERE ID IN (SELECT HCNID FROM Courses_HCN WHERE CourseID = ?)", courseID)
	rows, err := Db.Query("SELECT ID, CourseID, ClinicalCaseID, Displayable FROM Courses_CCases WHERE CourseID = ?", courseID)
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, CourseID, ClinicalCaseID, Displayable int
		if err := rows.Scan(&ID, &CourseID, &ClinicalCaseID, &Displayable); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var courseClinicalCase = mymodels.CourseClinicalCase{ID: &ID, CourseID: &CourseID, ClinicalCaseID: &ClinicalCaseID, Displayable: &Displayable}
		allCourseClinicalCase = append(allCourseClinicalCase, courseClinicalCase)
	}
	json.NewEncoder(w).Encode(allCourseClinicalCase)
	w.WriteHeader(http.StatusOK)
	return
}

// RemoveClinicalCase deletes a relationship between a course and a HCN...
func RemoveClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var removedCourseClinicalCase mymodels.CourseClinicalCase
	json.Unmarshal(reqBody, &removedCourseClinicalCase)

	switch {
	case (removedCourseClinicalCase.ClinicalCaseID) == nil || (*removedCourseClinicalCase.ClinicalCaseID*1 == 0) || (*removedCourseClinicalCase.ClinicalCaseID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (removedCourseClinicalCase.CourseID) == nil || (*removedCourseClinicalCase.CourseID*1 == 0) || (*removedCourseClinicalCase.CourseID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("DELETE FROM Courses_CCases WHERE CourseID=? AND ClinicalCaseID=?", removedCourseClinicalCase.CourseID, removedCourseClinicalCase.ClinicalCaseID)
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
}

// VisibilityClinicalCase bla bla...
func VisibilityClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedCourseClinicalCase mymodels.CourseClinicalCase
	json.Unmarshal(reqBody, &updatedCourseClinicalCase)
	switch {
	case (*updatedCourseClinicalCase.Displayable*1 > 1):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid display status", *updatedCourseClinicalCase.Displayable)
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Courses_CCases SET Displayable=? WHERE CourseID=? AND ClinicalCaseID=?", updatedCourseClinicalCase.Displayable, updatedCourseClinicalCase.CourseID, updatedCourseClinicalCase.ClinicalCaseID)
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
