package myhandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" //go get -u github.com/gorilla/mux
	"github.com/itrepablik/itrlog"
)

type productModel struct {
	Db *sql.DB
}

type student struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type allStudents []student

var students = allStudents{
	{
		ID:    1,
		Name:  "Xavier Garzón",
		Email: "xg@email.com",
	},
}

// GetStudents bla bla...
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students allStudents
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT Id, Name, Email FROM Students")
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		itrlog.Warn("(SQL) ", err.Error())
		return
	}
	for rows.Next() {
		var studentID int
		var Name, Email string
		if err := rows.Scan(&studentID, &Name, &Email); err != nil {
			itrlog.Warn("(SQL) ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		var student = student{ID: studentID, Name: Name, Email: Email}
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
	var Db, _ = config.MYSQLConnection()
	var Name, Email string
	err = Db.QueryRow("SELECT Name, Email FROM Students WHERE Id=?", studentID).Scan(&Name, &Email)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		itrlog.Warn("(SQL) ", err.Error())
		defer Db.Close()
		return
	}
	var student = student{ID: studentID, Name: Name, Email: Email}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
	w.WriteHeader(http.StatusCreated)
	return
}

// CreateStudent bla bla...
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newStudent)
	_, err = Db.Query("call createStudent(?, ?, ?)", newStudent.ID, newStudent.Name, newStudent.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		defer Db.Close()
		return
	}
	w.Header().Set("Content-Type", "application/json")
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

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("UPDATE Students SET Name=?, Email=? WHERE Id=?", updatedStudent.Name, updatedStudent.Email, updatedStudent.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		defer Db.Close()
		return
	}

	count, err := row.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		defer Db.Close()
		return
	}
	if count == 1 {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return
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
	_, err = Db.Exec("DELETE FROM Students WHERE Id=?", deletedStudent.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		defer Db.Close()
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return
}
