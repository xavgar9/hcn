package courses

import (
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	dbHelper "hcn/myhelpers/databaseHelper"
	stuctHelper "hcn/myhelpers/structValidationHelper"

	"github.com/itrepablik/sakto"
)

// CreateCourse creates one course in db.
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourse mymodels.Course
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newCourse)

	// Add creation date field
	CurrentLocalTime := sakto.GetCurDT(time.Now(), "America/New_York")
	NOW := strings.Split(CurrentLocalTime.String(), ".")[0]
	newCourse.CreationDate = &NOW

	// Fields validation
	structFields := []string{"TeacherID", "Name", "CreationDate"} // struct fields to check
	_, err = newCourse.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newCourse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

// GetAllCourses returns all courses in db.
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Data from db
	var course mymodels.Course
	rows, err := dbHelper.GetAll(course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var allCourses mymodels.AllCourses
	dbHelper.RowsToStruct(rows, &allCourses)

	json.NewEncoder(w).Encode(allCourses)
	return
}

// GetCourse returns one course filtered by the id.
func GetCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("ID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	course := mymodels.Course{ID: &id}

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = course.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from into db
	rows, err := dbHelper.Get(course)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allCourses mymodels.AllCourses
	dbHelper.RowsToStruct(rows, &allCourses)
	json.NewEncoder(w).Encode(allCourses[0])

	return
}

// UpdateCourse updates fields of an course in db.
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedCourse mymodels.Course
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedCourse)

	// Fields validation
	structFields := []string{"ID", "Name"} // struct fields to check
	_, err = updatedCourse.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedCourse)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// DeleteCourse deletes one course filtered by id.
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedCourse mymodels.Course
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedCourse)

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = deletedCourse.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedCourse)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}

	return
}

// AddStudent adds an Student into a course.
func AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudentTuition mymodels.StudentTuition
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newStudentTuition)

	// Fields validation
	structFields := []string{"CourseID", "StudentID"} // struct fields to check
	_, err = newStudentTuition.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newStudentTuition)
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

// GetAllStudentsCourseNoHTTP get all HCN in a course.
func GetAllStudentsCourseNoHTTP(courseID int) (mymodels.AllStudents, error) {
	var allStudents mymodels.AllStudents
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, Name, Email FROM Students WHERE ID IN (SELECT StudentID FROM Students_Courses WHERE CourseID = ?)", courseID)
	defer Db.Close()
	if err != nil {
		return allStudents, err
	}
	for rows.Next() {
		var ID int
		var Name, Email string
		if err := rows.Scan(&ID, &Name, &Email); err != nil {
			return allStudents, err
		}
		var student = mymodels.Student{ID: &ID, Name: &Name, Email: &Email}
		allStudents = append(allStudents, student)
	}
	return allStudents, nil
}

// GetAllStudentsCourse get all HCN in a course.
func GetAllStudentsCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("ID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	studentTuition := mymodels.StudentTuition{CourseID: &id}

	// Fields validation
	structFields := []string{"CourseID"} // struct fields to check
	_, err = studentTuition.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from db
	rows, err := dbHelper.GetAll(studentTuition)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var allStudentTuition mymodels.AllStudents
	dbHelper.RowsToStruct(rows, &allStudentTuition)

	if len(allStudentTuition) != 0 {
		json.NewEncoder(w).Encode(allStudentTuition)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "(db 2) element does not exist in db")
	}
	return
}

// RemoveStudent from a course.
func RemoveStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedStudentTuition mymodels.StudentTuition
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedStudentTuition)

	// Fields validation
	structFields := []string{"CourseID", "StudentID"} // struct fields to check
	_, err = deletedStudentTuition.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedStudentTuition)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}

	return
}
