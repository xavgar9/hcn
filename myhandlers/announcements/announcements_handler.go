package announcements

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

// CreateAnnouncement bla bla...
func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAnnouncement mymodels.Announcement
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newAnnouncement)
	switch {
	case (newAnnouncement.CourseID == nil) || (*newAnnouncement.CourseID*1 == 0) || (*newAnnouncement.CourseID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (newAnnouncement.Title == nil) || (len(*newAnnouncement.Title) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case (newAnnouncement.Description == nil) || (len(*newAnnouncement.Description) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Announcements(CourseID,Title,Description,CreationDate) VALUES (?,?,?,NOW())", newAnnouncement.CourseID, newAnnouncement.Title, newAnnouncement.Description)
		defer Db.Close()
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 1 {
			int64ID, _ := rows.LastInsertId()
			intID := int(int64ID)
			newAnnouncement.ID = &intID
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newAnnouncement)
		}
		return
	}
}

// GetAllAnnouncements bla bla...
func GetAllAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var announcements mymodels.AllAnnouncements
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, CourseID, Title, Description, CreationDate FROM Announcements")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
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
		var announcement = mymodels.Announcement{ID: &ID, CourseID: &CourseID, Title: &Title, Description: &Description, CreationDate: &CreationDate}
		announcements = append(announcements, announcement)
	}
	json.NewEncoder(w).Encode(announcements)
	w.WriteHeader(http.StatusOK)
	return
}

// GetAnnouncement bla bla...
func GetAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	announcementID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	var ID, CourseID int
	var Title, Description, CreationDate string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID,CourseID,Title,Description,CreationDate FROM Announcements WHERE ID=?", announcementID).Scan(&ID, &CourseID, &Title, &Description, &CreationDate)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var announcement = mymodels.Announcement{ID: &ID, CourseID: &CourseID, Title: &Title, Description: &Description, CreationDate: &CreationDate}
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
	var updatedAnnouncement mymodels.Announcement
	json.Unmarshal(reqBody, &updatedAnnouncement)
	switch {
	case (updatedAnnouncement.ID == nil) || (*updatedAnnouncement.ID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedAnnouncement.Title == nil) || len(*updatedAnnouncement.Title) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Title is empty or not valid")
		return
	case (updatedAnnouncement.Description == nil) || len(*updatedAnnouncement.Description) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Description is empty or not valid")
		return
	case (updatedAnnouncement.CreationDate == nil) || len(*updatedAnnouncement.CreationDate) == 0:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CreationDate is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Announcements SET Title=?, Description=?, CreationDate=? WHERE ID=?", updatedAnnouncement.Title, updatedAnnouncement.Description, updatedAnnouncement.CreationDate, updatedAnnouncement.ID)
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
			json.NewEncoder(w).Encode(updatedAnnouncement)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
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
	var deletedAnnouncement mymodels.Announcement
	json.Unmarshal(reqBody, &deletedAnnouncement)

	if (deletedAnnouncement.ID) == nil || (*deletedAnnouncement.ID*1 <= 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Announcements WHERE ID=?", deletedAnnouncement.ID)
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
