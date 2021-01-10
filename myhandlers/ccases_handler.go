package myhandlers

import (
	"encoding/json"
	"fmt"
	"hcn/config"
	"io/ioutil"
	"net/http"
)

// CreateClinicalCase bla bla...
func CreateClinicalCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClinicalCase ClinicalCase
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newClinicalCase)
	switch {
	case (*newClinicalCase.TeachersId*1 == 0) || (*newClinicalCase.TeachersId*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid TeachersID", *newClinicalCase.TeachersId)
		return
	case (newClinicalCase.Title == nil) || (len(*newClinicalCase.Title) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case (newClinicalCase.Description == nil) || (len(*newClinicalCase.Description) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case (newClinicalCase.Media == nil) || (len(*newClinicalCase.Media) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Media is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Clinical_Cases(Title,Description,Media,TeachersId) VALUES (?,?,?,?)", newClinicalCase.Title, newClinicalCase.Description, newClinicalCase.Media, newClinicalCase.TeachersId)
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

/*
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

// GetAnnouncement bla bla...
func GetAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	announcementID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var ID, CoursesID int
	var Title, Description, CreationDate string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID,CoursesID,Title,Description,CreationDate FROM Announcements WHERE Id=?", announcementID).Scan(&ID, &CoursesID, &Title, &Description, &CreationDate)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var announcement = Announcement{ID: &ID, CoursesID: &CoursesID, Title: &Title, Description: &Description, CreationDate: &CreationDate}
	json.NewEncoder(w).Encode(announcement)
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
	var updatedAnnouncement Announcement
	json.Unmarshal(reqBody, &updatedAnnouncement)
	switch {
	case (updatedAnnouncement.ID) == nil || (*updatedAnnouncement.ID*1 == 0) || (*updatedAnnouncement.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case updatedAnnouncement.Title == nil || len(*updatedAnnouncement.Title) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case updatedAnnouncement.Description == nil || len(*updatedAnnouncement.Description) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case updatedAnnouncement.CreationDate == nil || len(*updatedAnnouncement.CreationDate) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CreationDate is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Announcements SET Title=?, Description=?, CreationDate=? WHERE Id=?", updatedAnnouncement.Title, updatedAnnouncement.Description, updatedAnnouncement.CreationDate, updatedAnnouncement.ID)
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
	var deletedAnnouncement Announcement
	json.Unmarshal(reqBody, &deletedAnnouncement)

	if (deletedAnnouncement.ID) == nil || (*deletedAnnouncement.ID*1 == 0) || (*deletedAnnouncement.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Announcements WHERE Id=?", deletedAnnouncement.ID)
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
