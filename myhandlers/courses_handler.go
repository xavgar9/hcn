package myhandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/helpers"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateCourse bla bla...
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourse Course
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newCourse)
	switch {
	case (*newCourse.TeacherID*1 == 0) || (*newCourse.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Teacher is empty or not valid")
		return
	case (newCourse.Name == nil) || (len(*newCourse.Name) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Courses(TeacherID,Name,CreationDate) VALUES (?,?,NOW())", newCourse.TeacherID, newCourse.Name)
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

// GetAllCourses bla bla...
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courses AllCourses
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, TeacherID, Name, CreationDate FROM Courses")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, TeacherID int
		var Name, CreationDate string
		if err := rows.Scan(&ID, &TeacherID, &Name, &CreationDate); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var course = Course{ID: &ID, TeacherID: &TeacherID, Name: &Name, CreationDate: &CreationDate}
		courses = append(courses, course)
	}
	json.NewEncoder(w).Encode(courses)
	w.WriteHeader(http.StatusOK)
	return
}

// GetCourse bla bla...
func GetCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}
	var ID, TeacherID int
	var Name, CreationDate string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, TeacherID, Name, CreationDate FROM Courses WHERE ID=?", courseID).Scan(&ID, &TeacherID, &Name, &CreationDate)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var course = Course{ID: &ID, TeacherID: &TeacherID, Name: &Name, CreationDate: &CreationDate}
	json.NewEncoder(w).Encode(course)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateCourse bla bla...
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedCourse Course
	json.Unmarshal(reqBody, &updatedCourse)
	switch {
	case updatedCourse.Name == nil || len(*updatedCourse.Name) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Courses SET Name=? WHERE ID=?", updatedCourse.Name, updatedCourse.ID)
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

// DeleteCourse bla bla...
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedCourse Course
	json.Unmarshal(reqBody, &deletedCourse)

	if (deletedCourse.ID) == nil || (*deletedCourse.ID*1 == 0) || (*deletedCourse.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Courses WHERE ID=?", deletedCourse.ID)
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

// AddStudent adds an Student into a course...
func AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var newStudentTuition StudentTuition
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newStudentTuition)
	switch {
	case (newStudentTuition.CourseID == nil) || (*newStudentTuition.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CoursesID is empty or not valid")
		return
	case (newStudentTuition.StudentID == nil) || (*newStudentTuition.StudentID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "StudentID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Students_Courses(CourseID,StudentID) VALUES (?,?)", newStudentTuition.CourseID, newStudentTuition.StudentID)
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

// GetAllStudentsCourse get all HCN in a course...
func GetAllStudentsCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !helpers.VerifyRequest(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", vars["id"])
		return
	}

	var allStudents AllStudents
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, Name, Email FROM Students WHERE ID IN (SELECT StudentID FROM Students_Courses WHERE CourseID = ?)", courseID)
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID int
		var Name, Email string
		if err := rows.Scan(&ID, &Name, &Email); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var student = Student{ID: &ID, Name: &Name, Email: &Email}
		allStudents = append(allStudents, student)
	}
	json.NewEncoder(w).Encode(allStudents)
	w.WriteHeader(http.StatusOK)
	return
}

// RemoveStudent from a course...
func RemoveStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var removedStudent StudentTuition
	json.Unmarshal(reqBody, &removedStudent)
	if (removedStudent.StudentID) == nil || (*removedStudent.StudentID*1 == 0) || (*removedStudent.StudentID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "StudentID is empty or not valid")
		return
	}
	if (removedStudent.CourseID) == nil || (*removedStudent.CourseID*1 == 0) || (*removedStudent.CourseID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Students_Courses WHERE CourseID=? AND StudentID=?", removedStudent.CourseID, removedStudent.StudentID)
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
