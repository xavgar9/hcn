package teachers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"

	"strconv"
)

// CreateTeacher bla bla...
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTeacher mymodels.Teacher
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newTeacher)
	switch {
	case (newTeacher.ID == nil) || (*newTeacher.ID*1 == 0) || (*newTeacher.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (newTeacher.Name == nil) || (len(*newTeacher.Name) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	case (newTeacher.Email == nil) || (len(*newTeacher.Email) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Teachers(ID,Name,Email) VALUES (?, ?, ?)", newTeacher.ID, newTeacher.Name, newTeacher.Email)
		defer Db.Close()
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, err := rows.RowsAffected()
		if cnt == 1 {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTeacher)
		}
		return
	}
}

// GetAllTeachers bla bla...
func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var teachers mymodels.AllTeachers
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, Name, Email FROM Teachers")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var teacherID int
		var Name, Email string
		if err := rows.Scan(&teacherID, &Name, &Email); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var Teacher = mymodels.Teacher{ID: &teacherID, Name: &Name, Email: &Email}
		teachers = append(teachers, Teacher)
	}
	json.NewEncoder(w).Encode(teachers)
	w.WriteHeader(http.StatusOK)
	return
}

// GetTeacher bla bla...
func GetTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Url Param 'id' is missing or is invalid")
		return
	}
	teacherID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v is not a valid ID", keys[0])
		return
	}

	var Name, Email string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT Name, Email FROM Teachers WHERE ID=?", teacherID).Scan(&Name, &Email)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var Teacher = mymodels.Teacher{ID: &teacherID, Name: &Name, Email: &Email}
	json.NewEncoder(w).Encode(Teacher)
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
	var updatedTeacher mymodels.Teacher
	json.Unmarshal(reqBody, &updatedTeacher)
	switch {
	case (updatedTeacher.ID == nil) || (*updatedTeacher.ID*1 == 0) || (*updatedTeacher.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedTeacher.Name == nil) || (len(*updatedTeacher.Name) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty or not valid")
		return
	case (updatedTeacher.Email == nil) || (len(*updatedTeacher.Email) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Teachers SET Name=?, Email=? WHERE ID=?", updatedTeacher.Name, updatedTeacher.Email, updatedTeacher.ID)
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
			json.NewEncoder(w).Encode(updatedTeacher)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}

// DeleteTeacher bla bla...
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var deletedTeacher mymodels.Teacher
	json.Unmarshal(reqBody, &deletedTeacher)

	if (deletedTeacher.ID) == nil || (*deletedTeacher.ID*1 == 0) || (*deletedTeacher.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Teachers WHERE ID=?", deletedTeacher.ID)
	defer Db.Close()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
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
