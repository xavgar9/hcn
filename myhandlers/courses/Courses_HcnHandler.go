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
	case (*newCourseHCN.CourseID*1 == 0) || (*newCourseHCN.CourseID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (*newCourseHCN.HCNID*1 == 0) || (*newCourseHCN.HCNID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (*newCourseHCN.Displayable*1 == 0) || (*newCourseHCN.Displayable*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid display status", *newCourseHCN.Displayable)
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
	vars := mux.Vars(r)
	courseHCNID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var ID, CourseID, HCNID, Displayable int
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, CourseID, HCNID, Displayable FROM Courses_HCN WHERE ID=? AND CourseID", courseHCNID).Scan(&ID, &CourseID, &HCNID, &Displayable)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var courseHcn = mymodels.CourseHCN{ID: &ID, CourseID: &CourseID, HCNID: &HCNID, Displayable: &Displayable}
	json.NewEncoder(w).Encode(courseHcn)
	w.WriteHeader(http.StatusCreated)
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
	var deletedCourseHCN mymodels.CourseHCN
	json.Unmarshal(reqBody, &deletedCourseHCN)

	if (deletedCourseHCN.ID) == nil || (*deletedCourseHCN.ID*1 == 0) || (*deletedCourseHCN.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Courses_HCN WHERE ID=?", deletedCourseHCN.ID)
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
	case (*updatedCourseHCN.Displayable*1 <= 2):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid display status", *updatedCourseHCN.Displayable)
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Courses_HCN SET Displayable=? WHERE ID=?", updatedCourseHCN.Displayable, updatedCourseHCN.ID)
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
