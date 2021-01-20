package courses

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

// CreateCourse bla bla...
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourse mymodels.Course
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	err = json.Unmarshal(reqBody, &newCourse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	switch {
	case (newCourse.TeacherID == nil) || (*newCourse.TeacherID*1 == 0) || (*newCourse.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Teacher is empty or not valid")
		return
	case (newCourse.Name == nil) || (len(*newCourse.Name) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Courses(TeacherID,Name,CreationDate) VALUES (?,?,NOW())", newCourse.TeacherID, newCourse.Name)
		if err != nil {
			defer Db.Close()
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 1 {
			int64ID, _ := rows.LastInsertId()
			intID := int(int64ID)
			newCourse.ID = &intID
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCourse)
		}
		return
	}
}

// GetAllCourses bla bla...
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courses mymodels.AllCourses
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
		var course = mymodels.Course{ID: &ID, TeacherID: &TeacherID, Name: &Name, CreationDate: &CreationDate}
		courses = append(courses, course)
	}
	json.NewEncoder(w).Encode(courses)
	w.WriteHeader(http.StatusOK)
	return
}

// GetCourse bla bla...
func GetCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
		return
	}
	courseID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
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
	var course = mymodels.Course{ID: &ID, TeacherID: &TeacherID, Name: &Name, CreationDate: &CreationDate}
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
	var updatedCourse mymodels.Course
	json.Unmarshal(reqBody, &updatedCourse)
	switch {
	case (updatedCourse.ID == nil) || (*updatedCourse.ID*1 == 0) || (*updatedCourse.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedCourse.Name == nil) || len(*updatedCourse.Name) == 0:
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
			json.NewEncoder(w).Encode(updatedCourse)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
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
	var deletedCourse mymodels.Course
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
	return
}

// AddStudent adds an Student into a course...
func AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudentTuition mymodels.StudentTuition
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
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 0 {			
			fmt.FPrintf(w, "Student added")
		} else if cnt == 1 {
			fmt.FPrintf(w, "Student  not added")
		}
		return
	}
}

// GetAllStudentsCourse get all HCN in a course...
func GetAllStudentsCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
		return
	}
	courseID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
		return
	}

	var allStudents mymodels.AllStudents
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
		var student = mymodels.Student{ID: &ID, Name: &Name, Email: &Email}
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
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	var removedStudent mymodels.StudentTuition
	json.Unmarshal(reqBody, &removedStudent)
	if (removedStudent.StudentID) == nil || (*removedStudent.StudentID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "StudentID is empty or not valid")
		return
	}
	if (removedStudent.CourseID) == nil || (*removedStudent.CourseID*1 <= 0):
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
	return
}
