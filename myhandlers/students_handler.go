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
	json.NewEncoder(w).Encode(students)
}

// GetStudent bla bla...
func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, student := range students {
		if student.ID == studentID {
			json.NewEncoder(w).Encode(student)
		}
	}
}

// CreateStudent bla bla...
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buenas SQL")
	var newStudent student
	var Db, _ = config.MYSQLConnection()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert valid student")
	}
	json.Unmarshal(reqBody, &newStudent)
	_, err = Db.Query("call createStudent(?, ?, ?)", newStudent.ID, newStudent.Name, newStudent.Email)
	if err != nil {
		fmt.Fprintf(w, "Something wrong in MySQL")
	}
	defer Db.Close()
	return

	/*
		var newStudent student
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Insert valid student")
		}
		json.Unmarshal(reqBody, &newStudent)
		newStudent.ID = len(students) + 1
		students = append(students, newStudent)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)
	*/

}

// UpdateStudent bla bla...
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert valid student")
	}
	var updatedStudent student
	json.Unmarshal(reqBody, &updatedStudent)

	for i, student := range students {
		if student.ID == studentID {
			students[len(students)-1], students[i] = students[i], students[len(students)-1]
			students = students[:len(students)-1]
			updatedStudent.ID = student.ID
			students = append(students, updatedStudent)
			fmt.Fprintf(w, "The student %v was updated succesfully", studentID)
		}
	}
}

// DeleteStudent bla bla...
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, student := range students {
		if student.ID == studentID {
			students = append(students[:i], students[i+1:]...)
			fmt.Fprintf(w, "Student %v was deleted succesfully", studentID)
		}
	}
}
