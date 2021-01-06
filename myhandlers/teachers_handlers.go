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

type teacher struct {
	ID    *int    `json:"ID"`
	Name  *string `json:"Name"`
	Email *string `json:"Email"`
}

type allTeachers []teacher

// CreateTeacher bla bla...
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudent teacher
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
		_, err = Db.Query("call createTeacher(?, ?, ?)", newStudent.ID, newStudent.Name, newStudent.Email)
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

// GetTeachers bla bla...
func GetTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students allTeachers
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT Id, Name, Email FROM Teachers")
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
		var teacher = teacher{ID: &studentID, Name: &Name, Email: &Email}
		students = append(students, teacher)
	}
	json.NewEncoder(w).Encode(students)
	w.WriteHeader(http.StatusOK)
	return
}

// GetTeacher bla bla...
func GetTeacher(w http.ResponseWriter, r *http.Request) {
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
	err = Db.QueryRow("SELECT Name, Email FROM Teachers WHERE Id=?", studentID).Scan(&Name, &Email)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var teacher = teacher{ID: &studentID, Name: &Name, Email: &Email}
	json.NewEncoder(w).Encode(teacher)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateTeacher bla bla...
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedStudent teacher
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
		row, err := Db.Exec("UPDATE Teachers SET Name=?, Email=? WHERE Id=?", updatedStudent.Name, updatedStudent.Email, updatedStudent.ID)
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

// DeleteTeacher bla bla...
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedStudent teacher
	json.Unmarshal(reqBody, &deletedStudent)

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Teachers WHERE Id=?", deletedStudent.ID)
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
