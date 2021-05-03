package announcements

import (
	"encoding/json"
	"fmt"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	dbHelper "hcn/myhelpers/databaseHelper"
	stuctHelper "hcn/myhelpers/structValidationHelper"

	"github.com/itrepablik/sakto"
)

// CreateAnnouncement creates one announcement in db.
func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAnnouncement mymodels.Announcement
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newAnnouncement)

	// Add creation date field
	CurrentLocalTime := sakto.GetCurDT(time.Now(), "America/New_York")
	NOW := strings.Split(CurrentLocalTime.String(), ".")[0]
	newAnnouncement.CreationDate = &NOW

	// Fields validation
	structFields := []string{"CourseID", "Title", "Description", "CreationDate"} // struct fields to check
	_, err = newAnnouncement.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newAnnouncement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

// GetAllAnnouncements returns all announcements in db.
func GetAllAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Data from db
	var announcement mymodels.Announcement
	rows, err := dbHelper.GetAll(announcement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var allAnnouncements mymodels.AllAnnouncements
	dbHelper.RowsToStruct(rows, &allAnnouncements)

	json.NewEncoder(w).Encode(allAnnouncements)
	return
}

// GetAnnouncement returns one announcement filtered by the id.
func GetAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("ID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	announcement := mymodels.Announcement{ID: &id}

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = announcement.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from into db
	rows, err := dbHelper.Get(announcement)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allAnnouncements mymodels.AllAnnouncements
	dbHelper.RowsToStruct(rows, &allAnnouncements)
	json.NewEncoder(w).Encode(allAnnouncements[0])

	return
}

// UpdateAnnouncement updates fields of an announcement in db.
func UpdateAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedAnnouncement mymodels.Announcement
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedAnnouncement)

	// Fields validation
	_, err = updatedAnnouncement.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedAnnouncement)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// DeleteAnnouncement deletes one announcement filtered by the id.
func DeleteAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedAnnouncement mymodels.Announcement
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedAnnouncement)

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = deletedAnnouncement.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedAnnouncement)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}
