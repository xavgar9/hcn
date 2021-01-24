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
	case (newCourseHCN.Displayable == nil):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	default:
		if *newCourseHCN.Displayable == 0 || *newCourseHCN.Displayable == 1 {
			rows, err := Db.Exec("INSERT INTO Courses_HCN(CourseID,HCNID,Displayable) VALUES (?,?,?)", newCourseHCN.CourseID, newCourseHCN.HCNID, newCourseHCN.Displayable)
			defer Db.Close()
			if err != nil {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "(SQL) %v", err.Error())
				return
			}
			cnt, _ := rows.RowsAffected()
			if cnt == 1 {
				fmt.Fprintf(w, "HCN added to course")
			}
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	}
}

// GetAllHCNCourse bla bla...
func GetAllHCNCourse(w http.ResponseWriter, r *http.Request) {
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
	case (removedCourseHCN.CourseID == nil) || (*removedCourseHCN.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (removedCourseHCN.HCNID == nil) || (*removedCourseHCN.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("DELETE FROM Courses_HCN WHERE CourseID=? AND HCNID=?", removedCourseHCN.CourseID, removedCourseHCN.HCNID)
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
	case (updatedCourseHCN.ID == nil) || (*updatedCourseHCN.ID*1 == 0) || (*updatedCourseHCN.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedCourseHCN.CourseID == nil) || (*updatedCourseHCN.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (updatedCourseHCN.HCNID == nil) || (*updatedCourseHCN.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (updatedCourseHCN.Displayable == nil) || (*updatedCourseHCN.Displayable*1 > 1):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Courses_HCN SET Displayable=? WHERE ID=? AND CourseID=? AND HCNID=?", updatedCourseHCN.Displayable, updatedCourseHCN.ID, updatedCourseHCN.CourseID, updatedCourseHCN.HCNID)
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
			json.NewEncoder(w).Encode(updatedCourseHCN)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}
