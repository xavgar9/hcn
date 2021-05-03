package ccases

import (
	"encoding/json"
	"fmt"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strings"

	dbHelper "hcn/myhelpers/databaseHelper"
	stuctHelper "hcn/myhelpers/structValidationHelper"

	b64 "encoding/base64"
)

// CreateClinicalCase creates one clinical case in db.
func CreateClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClinicalCase mymodels.ClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newClinicalCase)

	// Fields validation
	structFields := []string{"Title", "Description", "Media", "TeacherID"} // struct fields to check
	_, err = newClinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Code media
	media := b64.StdEncoding.EncodeToString([]byte(*newClinicalCase.Media)) //convert to base64 (BLOB)
	newClinicalCase.Media = &media

	// Data insertion into db
	_, err = dbHelper.Insert(newClinicalCase)
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

// GetAllClinicalCases returns all clinical cases in db.
func GetAllClinicalCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Data from db
	var clinicalCase mymodels.ClinicalCase
	rows, err := dbHelper.GetAll(clinicalCase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var allClinicalCases mymodels.AllClinicalCases
	dbHelper.RowsToStruct(rows, &allClinicalCases)

	// Handle coded data
	for i, clinicalCase := range allClinicalCases {
		codedData := *clinicalCase.Media
		decodedData, err := b64.StdEncoding.DecodeString(codedData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "(Decoder) %v", err.Error())
			return
		}
		formatedDecodedData := string(decodedData)
		clinicalCase.Media = &formatedDecodedData
		allClinicalCases[i] = clinicalCase
	}

	json.NewEncoder(w).Encode(allClinicalCases)
	return
}

// GetClinicalCase returns one clinical case filtered by the id.
func GetClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("ID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	clinicalCase := mymodels.ClinicalCase{ID: &id}

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = clinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from into db
	rows, err := dbHelper.Get(clinicalCase)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allClinicalCases mymodels.AllClinicalCases
	dbHelper.RowsToStruct(rows, &allClinicalCases)

	// Handle coded data
	clinicalCase = allClinicalCases[0]

	codedData := *clinicalCase.Media
	decodedData, err := b64.StdEncoding.DecodeString(codedData)
	if err != nil {
		fmt.Fprintf(w, "(Decoder) %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	formatedDecodedData := string(decodedData)
	clinicalCase.Media = &formatedDecodedData

	json.NewEncoder(w).Encode(clinicalCase)

	return
}

// UpdateClinicalCase updates fields of a clinical case in db.
func UpdateClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedClinicalCase mymodels.ClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedClinicalCase)

	// Fields validation
	_, err = updatedClinicalCase.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Handle uncoded data
	codedData := b64.StdEncoding.EncodeToString([]byte(*updatedClinicalCase.Media)) //convert to base64 (BLOB)
	updatedClinicalCase.Media = &codedData

	// Data update into db
	_, err = dbHelper.Update(updatedClinicalCase)
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

// DeleteClinicalCase deletes one clinical case filtered by the id.
func DeleteClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedClinicalCase mymodels.ClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedClinicalCase)

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = deletedClinicalCase.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedClinicalCase)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// LinkHCN adds an HCN into a Clinical Case...
func LinkHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newVinculation mymodels.HCNVinculation
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newVinculation)

	// Fields validation
	structFields := []string{"ClinicalCaseID", "HCNID"} // struct fields to check
	_, err = newVinculation.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newVinculation)
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

// UnlinkHCN from a clinical case.
func UnlinkHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedVinculation mymodels.HCNVinculation
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedVinculation)

	// Fields validation
	structFields := []string{"ClinicalCaseID", "HCNID"} // struct fields to check
	_, err = deletedVinculation.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedVinculation)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}
