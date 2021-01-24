package courses

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
	case (newCourseClinicalCase.CourseID == nil) || (*newCourseClinicalCase.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (newCourseClinicalCase.ClinicalCaseID == nil) || (*newCourseClinicalCase.ClinicalCaseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (newCourseClinicalCase.Displayable == nil):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	default:
		if *newCourseClinicalCase.Displayable == 0 || *newCourseClinicalCase.Displayable == 1 {
			rows, err := Db.Exec("INSERT INTO Courses_CCases(CourseID,ClinicalCaseID,Displayable) VALUES (?,?,?)", newCourseClinicalCase.CourseID, newCourseClinicalCase.ClinicalCaseID, newCourseClinicalCase.Displayable)
			defer Db.Close()
			if err != nil {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "(SQL) %v", err.Error())
				return
			}
			cnt, _ := rows.RowsAffected()
			if cnt == 1 {
				fmt.Fprintf(w, "Clinical Case added to course")
			}
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	}
}

// GetAllClinicalCases bla bla...
func GetAllClinicalCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	courseID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var allCourseClinicalCase mymodels.AllCourseClinicalCase
	var Db, _ = config.MYSQLConnection()
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
	case (removedCourseClinicalCase.CourseID) == nil || (*removedCourseClinicalCase.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (removedCourseClinicalCase.ClinicalCaseID) == nil || (*removedCourseClinicalCase.ClinicalCaseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
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
	case (updatedCourseClinicalCase.ID == nil) || (*updatedCourseClinicalCase.ID*1 == 0) || (*updatedCourseClinicalCase.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedCourseClinicalCase.CourseID == nil) || (*updatedCourseClinicalCase.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (updatedCourseClinicalCase.ClinicalCaseID == nil) || (*updatedCourseClinicalCase.ClinicalCaseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (updatedCourseClinicalCase.Displayable == nil) || (*updatedCourseClinicalCase.Displayable*1 > 1):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Courses_CCases SET Displayable=? WHERE CourseID=? AND ClinicalCaseID=?", updatedCourseClinicalCase.Displayable, updatedCourseClinicalCase.CourseID, updatedCourseClinicalCase.ClinicalCaseID)
		defer Db.Close()
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}

		count, err := row.RowsAffected()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		if count == 1 {
			json.NewEncoder(w).Encode(updatedCourseClinicalCase)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}
