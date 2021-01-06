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
	"github.com/itrepablik/itrlog"
)

/*
type productModel struct {
	Db *sql.DB
}
*/

type student struct {
	ID    *int    `json:"ID"`
	Name  *string `json:"Name"`
	Email *string `json:"Email"`
}

type allStudents []student

// CreateStudent bla bla...
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudent student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newStudent)
	switch {
	case newStudent.ID == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty")
		return
	case newStudent.Name == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty")
		return
	case newStudent.Email == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is empty")
		return
	default:
		_, err = Db.Query("call createStudent(?, ?, ?)", newStudent.ID, newStudent.Name, newStudent.Email)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		json.NewEncoder(w).Encode(newStudent)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// GetStudents bla bla...
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students allStudents
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT Id, Name, Email FROM Students")
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
		var student = student{ID: &studentID, Name: &Name, Email: &Email}
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)
	w.WriteHeader(http.StatusOK)
	return
}

// GetStudent bla bla...
func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		itrlog.Warn("(USER) ", err.Error())
		return
	}
	var Name, Email string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT Name, Email FROM Students WHERE Id=?", studentID).Scan(&Name, &Email)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var student = student{ID: &studentID, Name: &Name, Email: &Email}
	json.NewEncoder(w).Encode(student)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateStudent bla bla...
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedStudent student
	json.Unmarshal(reqBody, &updatedStudent)
	switch {
	case updatedStudent.ID == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty")
		return
	case updatedStudent.Name == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty")
		return
	case updatedStudent.Email == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is empty")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Students SET Name=?, Email=? WHERE Id=?", updatedStudent.Name, updatedStudent.Email, updatedStudent.ID)
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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedStudent student
	json.Unmarshal(reqBody, &deletedStudent)

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Students WHERE Id=?", deletedStudent.ID)
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
