package activities

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

// CreateActivity bla bla...
func CreateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newActivity mymodels.Activity
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newActivity)
	switch {
	case (newActivity.Title == nil) || (len(*newActivity.Title) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case (newActivity.Description == nil) || (len(*newActivity.Description) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case (newActivity.Type == nil) || (len(*newActivity.Type) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Type is empty or not valid")
		return
	case (newActivity.LimitDate == nil) || (len(*newActivity.LimitDate) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "LimitDate is empty or not valid")
		return
	case (newActivity.CourseID == nil) || (*newActivity.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (newActivity.ClinicalCaseID == nil) || (*newActivity.ClinicalCaseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (newActivity.HCNID == nil) || (*newActivity.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (newActivity.Difficulty == nil) || (*newActivity.Difficulty*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Difficulty is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Activities(Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty) VALUES (?,?,?,NOW(),?,?,?,?,?)", newActivity.Title, newActivity.Description, newActivity.Type, newActivity.LimitDate, newActivity.CourseID, newActivity.ClinicalCaseID, newActivity.HCNID, newActivity.Difficulty)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 1 {
			int64ID, _ := rows.LastInsertId()
			intID := int(int64ID)
			newActivity.ID = &intID
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newActivity)
		}
		return
	}
}

// GetAllActivities bla bla...
func GetAllActivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activities mymodels.AllActivities
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID,Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty FROM Activities")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, CoursesID, ClinicalCaseID, HCNID, Difficulty int
		var Title, Description, Type, CreationDate, LimitDate string
		if err := rows.Scan(&ID, &Title, &Description, &Type, &CreationDate, &LimitDate, &CoursesID, &ClinicalCaseID, &HCNID, &Difficulty); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var activity = mymodels.Activity{ID: &ID, Title: &Title, Description: &Description, Type: &Type, CreationDate: &CreationDate, LimitDate: &LimitDate, CourseID: &CoursesID, ClinicalCaseID: &ClinicalCaseID, HCNID: &HCNID, Difficulty: &Difficulty}
		activities = append(activities, activity)
	}
	json.NewEncoder(w).Encode(activities)
	w.WriteHeader(http.StatusOK)
	return
}

// GetActivity bla bla...
func GetActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	activityID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	var ID, CourseID, ClinicalCaseID, HCNID, Difficulty int
	var Title, Description, Type, CreationDate, LimitDate string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID,Title,Description,Type,CreationDate,LimitDate,CourseID,ClinicalCaseID,HCNID,Difficulty FROM Activities WHERE ID=?", activityID).Scan(&ID, &Title, &Description, &Type, &CreationDate, &LimitDate, &CourseID, &ClinicalCaseID, &HCNID, &Difficulty)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
		}
		return
	}
	var activity = mymodels.Activity{ID: &ID, Title: &Title, Description: &Description, Type: &Type, CreationDate: &CreationDate, LimitDate: &LimitDate, CourseID: &CourseID, ClinicalCaseID: &ClinicalCaseID, HCNID: &HCNID, Difficulty: &Difficulty}
	json.NewEncoder(w).Encode(activity)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateActivity bla bla...
func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedActivity mymodels.Activity
	json.Unmarshal(reqBody, &updatedActivity)
	switch {
	case (updatedActivity.ID == nil) || (*updatedActivity.ID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedActivity.Title == nil) || (len(*updatedActivity.Title) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case (updatedActivity.Description == nil) || (len(*updatedActivity.Description) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case (updatedActivity.Type == nil) || (len(*updatedActivity.Type) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Type is empty or not valid")
		return
	case (updatedActivity.LimitDate == nil) || (len(*updatedActivity.LimitDate) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "LimitDate is empty or not valid")
		return
	case (updatedActivity.ClinicalCaseID == nil) || (*updatedActivity.ClinicalCaseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ClinicalCaseID is empty or not valid")
		return
	case (updatedActivity.HCNID == nil) || (*updatedActivity.HCNID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "HCNID is empty or not valid")
		return
	case (updatedActivity.Difficulty == nil) || (*updatedActivity.Difficulty*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Difficulty is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Activities SET Title=?, Description=?, Type=?, LimitDate=?, ClinicalCaseID=?, HCNID=?, Difficulty=? WHERE ID=?", updatedActivity.Title, updatedActivity.Description, updatedActivity.Type, updatedActivity.LimitDate, updatedActivity.ClinicalCaseID, updatedActivity.HCNID, updatedActivity.Difficulty, updatedActivity.ID)
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
			json.NewEncoder(w).Encode(updatedActivity)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}

// DeleteActivity bla bla...
func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedActivity mymodels.Activity
	json.Unmarshal(reqBody, &deletedActivity)

	if (deletedActivity.ID) == nil || (*deletedActivity.ID*1 == 0) || (*deletedActivity.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Activities WHERE ID=?", deletedActivity.ID)
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
