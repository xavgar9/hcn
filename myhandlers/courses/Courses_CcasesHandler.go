package courses

import (
	"encoding/json"
	"fmt"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strings"

	dbHelper "hcn/myhelpers/databaseHelper"
	stuctHelper "hcn/myhelpers/structValidationHelper"
)

// AddClinicalCase adds a relationship between a course and a clinical case.
func AddClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourseClinicalCase mymodels.CourseClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newCourseClinicalCase)

	// Fields validation
	structFields := []string{"ClinicalCaseID", "CourseID", "Displayable"} // struct fields to check
	_, err = newCourseClinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newCourseClinicalCase)
	if err != nil {
		if strings.Split(err.Error(), ":")[0] == "(db 2) Error 1062" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

	return
}

// GetAllClinicalCases returns all clinical cases filtered by id.
func GetAllClinicalCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("CourseID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	courseClinicalCase := mymodels.CourseClinicalCase{CourseID: &id}

	// Fields validation
	structFields := []string{"CourseID"} // struct fields to check
	_, err = courseClinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from db
	rows, err := dbHelper.Get(courseClinicalCase)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allCourseClinicalCase mymodels.AllCourseClinicalCase
	dbHelper.RowsToStruct(rows, &allCourseClinicalCase)

	json.NewEncoder(w).Encode(allCourseClinicalCase)

	return
}

// RemoveClinicalCase deletes a relationship between a course and a HCN...
func RemoveClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedCourseClinicalCase mymodels.CourseClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedCourseClinicalCase)

	// Fields validation
	structFields := []string{"CourseID", "ClinicalCaseID"} // struct fields to check
	_, err = deletedCourseClinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedCourseClinicalCase)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// VisibilityClinicalCase updates the visibility of an clinical case.
func VisibilityClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedCourseClinicalCase mymodels.CourseClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedCourseClinicalCase)

	// Fields validation
	structFields := []string{"CourseID", "ClinicalCaseID", "Displayable"} // struct fields to check
	_, err = updatedCourseClinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedCourseClinicalCase)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, err.Error())
	}

	return
}
