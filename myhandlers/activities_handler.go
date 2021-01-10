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

// CreateActivity bla bla...
func CreateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newActivity Activity
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
	case (*newActivity.CoursesID*1 == 0) || (*newActivity.CoursesID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid CourseID", *newActivity.CoursesID)
		return
	case (*newActivity.ClinicalCasesID*1 == 0) || (*newActivity.ClinicalCasesID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid ClinicalCasesID", *newActivity.CoursesID)
		return
	case (*newActivity.Difficulty*1 == 0) || (*newActivity.Difficulty*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid Difficulty", *newActivity.CoursesID)
		return
	default:
		rows, err := Db.Exec("INSERT INTO Activities(Title,Description,Type,CreationDate,LimitDate,CoursesId,ClinicalCasesId,Difficulty) VALUES (?,?,?,NOW(),?,?,?,?)", newActivity.Title, newActivity.Description, newActivity.Type, newActivity.LimitDate, newActivity.CoursesID, newActivity.ClinicalCasesID, newActivity.Difficulty)
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

// GetActivities bla bla...
func GetActivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activities AllActivities
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT Id,Title,Description,Type,CreationDate,LimitDate,CoursesId,ClinicalCasesId,Difficulty FROM Activities")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, CoursesID, ClinicalCasesID, Difficulty int
		var Title, Description, Type, CreationDate, LimitDate string
		if err := rows.Scan(&ID, &Title, &Description, &Type, &CreationDate, &LimitDate, &CoursesID, &ClinicalCasesID, &Difficulty); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var activity = Activity{ID: &ID, Title: &Title, Description: &Description, Type: &Type, CreationDate: &CreationDate, LimitDate: &LimitDate, CoursesID: &CoursesID, ClinicalCasesID: &ClinicalCasesID, Difficulty: &Difficulty}
		activities = append(activities, activity)
	}
	json.NewEncoder(w).Encode(activities)
	w.WriteHeader(http.StatusOK)
	return
}

// GetActivity bla bla...
func GetActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	activityID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var ID, CoursesID, ClinicalCasesID, Difficulty int
	var Title, Description, Type, CreationDate, LimitDate string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT Id,Title,Description,Type,CreationDate,LimitDate,CoursesId,ClinicalCasesId,Difficulty FROM Activities WHERE Id=?", activityID).Scan(&ID, &Title, &Description, &Type, &CreationDate, &LimitDate, &CoursesID, &ClinicalCasesID, &Difficulty)
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
	var activity = Activity{ID: &ID, Title: &Title, Description: &Description, Type: &Type, CreationDate: &CreationDate, LimitDate: &LimitDate, CoursesID: &CoursesID, ClinicalCasesID: &ClinicalCasesID, Difficulty: &Difficulty}
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
	var updatedActivity Activity
	json.Unmarshal(reqBody, &updatedActivity)
	switch {
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
	case (*updatedActivity.ClinicalCasesID*1 == 0) || (*updatedActivity.ClinicalCasesID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid ClinicalCasesID", *updatedActivity.CoursesID)
		return
	case (*updatedActivity.Difficulty*1 == 0) || (*updatedActivity.Difficulty*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid Difficulty", *updatedActivity.CoursesID)
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Activities SET Title=?, Description=?, Type=?, LimitDate=?, ClinicalCasesId=?, Difficulty=? WHERE Id=?", updatedActivity.Title, updatedActivity.Description, updatedActivity.Type, updatedActivity.LimitDate, updatedActivity.ClinicalCasesID, updatedActivity.Difficulty, updatedActivity.ID)
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

// DeleteActivity bla bla...
func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedActivity Activity
	json.Unmarshal(reqBody, &deletedActivity)

	if (deletedActivity.ID) == nil || (*deletedActivity.ID*1 == 0) || (*deletedActivity.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Activities WHERE Id=?", deletedActivity.ID)
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
