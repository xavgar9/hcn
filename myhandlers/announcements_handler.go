package myhandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"io/ioutil"
	"net/http"
)

// CreateAnnouncement bla bla...
func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAnnouncement Announcement
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newAnnouncement)
	switch {
	case (*newAnnouncement.CoursesID*1 == 0) || (*newAnnouncement.CoursesID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid CourseID", *newAnnouncement.CoursesID)
		return
	case newAnnouncement.Title == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty")
		return
	case newAnnouncement.Description == nil:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty")
		return
	default:
		rows, err := Db.Query("INSERT INTO Announcements(CoursesID,Title,Description,CreationDate) VALUES (?,?,?,?,NOW())", newAnnouncement.CoursesID, newAnnouncement.Title, newAnnouncement.Description)
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
			json.NewEncoder(w).Encode(newAnnouncement)
		}
		return
	}
}

// GetAnnouncements bla bla...
func GetAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var announcements AllAnnouncements
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT Id, CoursesID, Title, Description, CreationDate FROM Announcements")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, CourseID int
		var Title, Description, CreationDate string
		if err := rows.Scan(&ID, &CourseID, &Title, &Description, &CreationDate); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var announcement = Announcement{ID: &ID, CoursesID: &CourseID, Title: &Title, Description: &Description, CreationDate: &CreationDate}
		announcements = append(announcements, announcement)
	}
	json.NewEncoder(w).Encode(announcements)
	w.WriteHeader(http.StatusOK)
	return
}

/*
// GetAnnouncement bla bla...
func GetAnnouncement(w http.ResponseWriter, r *http.Request) {
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
	err = Db.QueryRow("SELECT Name, Email FROM Announcements WHERE Id=?", studentID).Scan(&Name, &Email)
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

// UpdateAnnouncement bla bla...
func UpdateAnnouncement(w http.ResponseWriter, r *http.Request) {
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
		row, err := Db.Exec("UPDATE Announcements SET Name=?, Email=? WHERE Id=?", updatedStudent.Name, updatedStudent.Email, updatedStudent.ID)
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

// DeleteAnnouncement bla bla...
func DeleteAnnouncement(w http.ResponseWriter, r *http.Request) {
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
	row, err := Db.Exec("DELETE FROM Announcements WHERE Id=?", deletedStudent.ID)
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
*/
