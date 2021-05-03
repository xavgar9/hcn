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

// AddHCN adds a relationship between a course and a HCN.
func AddHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourseHCN mymodels.CourseHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newCourseHCN)

	// Fields validation
	structFields := []string{"HCNID", "CourseID", "Displayable"} // struct fields to check
	_, err = newCourseHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newCourseHCN)
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

// GetAllHCNCourse returns all hcns filtered by id.
func GetAllHCNCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("CourseID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	courseCourseHCN := mymodels.CourseHCN{CourseID: &id}

	// Fields validation
	structFields := []string{"CourseID"} // struct fields to check
	_, err = courseCourseHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from db
	rows, err := dbHelper.Get(courseCourseHCN)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allCourseCourseHCN mymodels.AllCourseHCN
	dbHelper.RowsToStruct(rows, &allCourseCourseHCN)

	json.NewEncoder(w).Encode(allCourseCourseHCN)

	return
}

// RemoveHCN deletes a relationship between a course and a HCN.
func RemoveHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedCourseHCN mymodels.CourseHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedCourseHCN)

	// Fields validation
	structFields := []string{"CourseID", "HCNID"} // struct fields to check
	_, err = deletedCourseHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedCourseHCN)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// VisibilityHCN updates the visibility of an hcn.
func VisibilityHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedCourseHCN mymodels.CourseHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedCourseHCN)

	// Fields validation
	structFields := []string{"CourseID", "HCNID", "Displayable"} // struct fields to check
	_, err = updatedCourseHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedCourseHCN)
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
