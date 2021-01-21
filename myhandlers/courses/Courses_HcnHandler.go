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

// AddHCN adds a relationship between a course and a HCN...
func AddHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourseHCN mymodels.CourseHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newCourseHCN)
	switch {
	case (newCourseHCN.CourseID == nil) || (*newCourseHCN.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (newCourseHCN.HCNID == nil) || (*newCourseHCN.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (newCourseHCN.Displayable == nil) || (*newCourseHCN.Displayable*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable value is not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Courses_HCN(CourseID,HCNID,Displayable) VALUES (?,?,?)", newCourseHCN.CourseID, newCourseHCN.HCNID, newCourseHCN.Displayable)
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

// GetAllHCN bla bla...
func GetAllHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var allCourseHCN mymodels.AllCourseHCN
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, CourseID, HCNID, Displayable FROM Courses_HCN WHERE CourseID = ?", courseID)
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, CourseID, HCNID, Displayable int
		if err := rows.Scan(&ID, &CourseID, &HCNID, &Displayable); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var courseHCN = mymodels.CourseHCN{ID: &ID, CourseID: &CourseID, HCNID: &HCNID, Displayable: &Displayable}
		allCourseHCN = append(allCourseHCN, courseHCN)
	}
	json.NewEncoder(w).Encode(allCourseHCN)
	w.WriteHeader(http.StatusOK)
	return
}

// RemoveHCN deletes a relationship between a course and a HCN...
func RemoveHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var removedCourseHCN mymodels.CourseHCN
	json.Unmarshal(reqBody, &removedCourseHCN)

	switch {
	case (removedCourseHCN.HCNID) == nil || (*removedCourseHCN.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (removedCourseHCN.CourseID) == nil || (*removedCourseHCN.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("DELETE FROM Courses_HCN WHERE CourseID=? AND HCNID=?", removedCourseHCN.CourseID, removedCourseHCN.HCNID)
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

// VisibilityHCN bla bla...
func VisibilityHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedCourseHCN mymodels.CourseHCN
	json.Unmarshal(reqBody, &updatedCourseHCN)
	switch {
	case (updatedCourseHCN.Displayable == nil) || (*updatedCourseHCN.Displayable*1 > 1):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable value is not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Courses_HCN SET Displayable=? WHERE CourseID=? AND HCNID=?", updatedCourseHCN.Displayable, updatedCourseHCN.CourseID, updatedCourseHCN.HCNID)
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
		return
	}
}
