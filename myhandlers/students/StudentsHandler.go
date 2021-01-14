package students

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

// CreateStudent bla bla...
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var newStudent mymodels.Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newStudent)
	switch {
	case (newStudent.ID == nil) || (*newStudent.ID*1 == 0) || (*newStudent.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (newStudent.Name == nil) || (len(*newStudent.Name) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	case (newStudent.Email == nil) || (len(*newStudent.Email) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is empty or not valid")
		return
	default:
		rows, err := Db.Query("INSERT INTO Students(ID,Name,Email) VALUES (?, ?, ?)", newStudent.ID, newStudent.Name, newStudent.Email)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt := 0
		for rows.Next() {
			cnt++
		}
		if cnt == 1 {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newStudent)
		}
		return
	}
}

// GetAllStudents bla bla...
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var Students mymodels.AllStudents
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, Name, Email FROM Students")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var studentID int
		var Name, Email string
		if err := rows.Scan(&studentID, &Name, &Email); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var Student = mymodels.Student{ID: &studentID, Name: &Name, Email: &Email}
		Students = append(Students, Student)
	}
	json.NewEncoder(w).Encode(Students)
	w.WriteHeader(http.StatusOK)
	return
}

// GetStudent bla bla...
func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var Name, Email string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT Name, Email FROM Students WHERE ID=?", studentID).Scan(&Name, &Email)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var Student = mymodels.Student{ID: &studentID, Name: &Name, Email: &Email}
	json.NewEncoder(w).Encode(Student)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateStudent bla bla...
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedStudent mymodels.Student
	json.Unmarshal(reqBody, &updatedStudent)
	switch {
	case updatedStudent.ID == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case updatedStudent.Name == nil || len(*updatedStudent.Name) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	case updatedStudent.Email == nil || len(*updatedStudent.Email) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Students SET Name=?, Email=? WHERE ID=?", updatedStudent.Name, updatedStudent.Email, updatedStudent.ID)
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

// DeleteStudent bla bla...
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedStudent mymodels.Student
	json.Unmarshal(reqBody, &deletedStudent)

	if (deletedStudent.ID) == nil || (*deletedStudent.ID*1 == 0) || (*deletedStudent.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Students WHERE ID=?", deletedStudent.ID)
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
