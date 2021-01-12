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

// CreateHCN bla bla...
func CreateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newHCN HCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newHCN)
	switch {
	case (*newHCN.TeacherID*1 == 0) || (*newHCN.TeacherID*1 < 0):
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
	var hcns AllHCN
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
		var hcn = HCN{ID: &ID, TeacherID: &TeacherID}
		hcns = append(hcns, hcn)
	}
	json.NewEncoder(w).Encode(hcns)
	w.WriteHeader(http.StatusOK)
	return
}

// GetHCN bla bla...
func GetHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	hcnID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
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
	var course = HCN{ID: &ID, TeacherID: &TeacherID}
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
	var updatedHCN HCN
	json.Unmarshal(reqBody, &updatedHCN)
	switch {
	case (*updatedHCN.TeacherID*1 == 0) || (*updatedHCN.TeacherID*1 < 0):
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
			fmt.Fprintf(w, "One row updated")
		} else {
			fmt.Fprintf(w, "No rows updated")
		}

		w.Header().Set("Content-Type", "application/json")
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
	var deletedHCN HCN
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

// ---------------------------------------------------------------------
// HCN_Course is a table that represents a relationship
// between a HCN and a Course.
// ---------------------------------------------------------------------

/*

// GetHCNCourse bla bla...
func GetHCNCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	courseHCNID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var ID, CoursesID, HCNID, Displayable int
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, CoursesID, HCNID, Displayable FROM HCN_Courses WHERE ID=?", courseHCNID).Scan(&ID, &CoursesID, &HCNID, &Displayable)
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
	var courseHcn = HCNCourse{ID: &ID, CoursesID: &CoursesID, HCNID: &HCNID, Displayable: &Displayable}
	json.NewEncoder(w).Encode(courseHcn)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateHCNCourse bla bla...
func UpdateHCNCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedCourseHCN HCNCourse
	json.Unmarshal(reqBody, &updatedCourseHCN)
	switch {
	case (updatedCourseHCN.CoursesID == nil) || (*updatedCourseHCN.CoursesID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CoursesID is empty or not valid")
		return
	case (updatedCourseHCN.HCNID == nil) || (*updatedCourseHCN.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (updatedCourseHCN.Displayable == nil):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Displayable is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE HCN_Courses SET CoursesID=?, HCNID=?, Displayable=? WHERE ID=?", updatedCourseHCN.CoursesID, updatedCourseHCN.HCNID, updatedCourseHCN.Displayable, updatedCourseHCN.ID)
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

// DeleteHCNCourse deletes a relationship between a HCN and a course...
func DeleteHCNCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedCourseHCN HCNCourse
	json.Unmarshal(reqBody, &deletedCourseHCN)

	if (deletedCourseHCN.ID) == nil || (*deletedCourseHCN.ID*1 == 0) || (*deletedCourseHCN.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM HCN_Courses WHERE ID=?", deletedCourseHCN.ID)
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

// ---------------------------------------------------------------------
// HCN_CCases is a table that represents a relationship
// between a HCN and a Clinical Case.
// ---------------------------------------------------------------------

// CreateHCNCCase creates a relationship between a HCN and a course...
func CreateHCNCCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newHCNCCase HCNCCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newHCNCCase)
	switch {
	case (newHCNCCase.ClinicalCasesID == nil) || (*newHCNCCase.ClinicalCasesID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CCasesId is empty or not valid")
		return
	case (newHCNCCase.HCNID == nil) || (*newHCNCCase.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO HCN_CCases(ClinicalCasesID,HCNID) VALUES (?,?)", newHCNCCase.ClinicalCasesID, newHCNCCase.HCNID)
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

// GetHCNsCCases bla bla...
func GetHCNsCCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hcnsCCases AllHCNsCCases
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, ClinicalCasesID, HCNID FROM HCN_CCases")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, ClinicalCasesID, HCNID int
		if err := rows.Scan(&ID, &ClinicalCasesID, &HCNID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var hcnCCase = HCNCCase{ID: &ID, ClinicalCasesID: &ClinicalCasesID, HCNID: &HCNID}
		hcnsCCases = append(hcnsCCases, hcnCCase)
	}
	json.NewEncoder(w).Encode(hcnsCCases)
	w.WriteHeader(http.StatusOK)
	return
}

// GetHCNCCase bla bla...
func GetHCNCCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	HCNCCaseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var ID, ClinicalCasesID, HCNID int
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, ClinicalCasesID, HCNID FROM HCN_CCases WHERE ID=?", HCNCCaseID).Scan(&ID, &ClinicalCasesID, &HCNID)
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
	var hcnCCAse = HCNCCase{ID: &ID, ClinicalCasesID: &ClinicalCasesID, HCNID: &HCNID}
	json.NewEncoder(w).Encode(hcnCCAse)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateHCNCCase bla bla...
func UpdateHCNCCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedHCNCCase HCNCCase
	json.Unmarshal(reqBody, &updatedHCNCCase)
	switch {
	case (updatedHCNCCase.ClinicalCasesID == nil) || (*updatedHCNCCase.ClinicalCasesID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCasesID is empty or not valid")
		return
	case (updatedHCNCCase.HCNID == nil) || (*updatedHCNCCase.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE HCN_CCases SET ClinicalCasesID=?, HCNID=? WHERE ID=?", updatedHCNCCase.ClinicalCasesID, updatedHCNCCase.HCNID, updatedHCNCCase.ID)
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

// DeleteHCNCCase deletes a relationship between a HCN and a course...
func DeleteHCNCCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedHCNCCase HCNCCase
	json.Unmarshal(reqBody, &deletedHCNCCase)

	if (deletedHCNCCase.ID) == nil || (*deletedHCNCCase.ID*1 == 0) || (*deletedHCNCCase.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM HCN_CCases WHERE ID=?", deletedHCNCCase.ID)
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
*/
