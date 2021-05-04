package activities

import (
	"encoding/json"
	"fmt"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	dbHelper "hcn/myhelpers/databaseHelper"
	stuctHelper "hcn/myhelpers/structValidationHelper"

	"github.com/itrepablik/sakto"
)

// CreateActivity creates one activity in db.
func CreateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newActivity mymodels.Activity
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newActivity)

	// Add creation date field
	CurrentLocalTime := sakto.GetCurDT(time.Now(), "America/New_York")
	NOW := strings.Split(CurrentLocalTime.String(), ".")[0]
	newActivity.CreationDate = &NOW

	// Fields validation
	structFields := []string{"Title", "Description", "Type", "LimitDate", "CourseID", "ClinicalCaseID", "HCNID", "Difficulty"} // struct fields to check
	_, err = newActivity.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	lastID, err := dbHelper.Insert(newActivity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	newActivity.ID = &lastID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newActivity)
	return
}

// GetAllActivities returns all activities in db.
func GetAllActivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Data from db
	var activity mymodels.Activity
	rows, err := dbHelper.GetAll(activity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var allActivities mymodels.AllActivities
	dbHelper.RowsToStruct(rows, &allActivities)

	json.NewEncoder(w).Encode(allActivities)
	return
}

// GetActivity returns one activity filtered by the id.
func GetActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("ID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	activity := mymodels.Activity{ID: &id}

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = activity.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from into db
	rows, err := dbHelper.Get(activity)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allActivities mymodels.AllActivities
	dbHelper.RowsToStruct(rows, &allActivities)
	json.NewEncoder(w).Encode(allActivities[0])

	return
}

// UpdateActivity updates fields of an announcement in db.
func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedActivity mymodels.Activity
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedActivity)

	// Fields validation
	_, err = updatedActivity.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedActivity)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// DeleteActivity deletes one activity filtered by id.
func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedActivity mymodels.Activity
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedActivity)

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = deletedActivity.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedActivity)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}

	return
}
