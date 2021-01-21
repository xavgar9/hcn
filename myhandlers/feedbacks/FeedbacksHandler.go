package feedbacks

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

// CreateFeedback bla bla...
func CreateFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newFeedback mymodels.Feedback
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	err = json.Unmarshal(reqBody, &newFeedback)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	switch {
	case (newFeedback.ActivityID == nil) || (*newFeedback.ActivityID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ActivityID is empty or not valid")
		return
	case (newFeedback.StudentID == nil) || (*newFeedback.StudentID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "StudentID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO Feedbacks(ActivityID,StudentID) VALUES (?,?)", newFeedback.ActivityID, newFeedback.StudentID)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 1 {
			int64ID, _ := rows.LastInsertId()
			intID := int(int64ID)
			newFeedback.ID = &intID
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newFeedback)
		}
		return
	}
}

// GetAllFeedbacks bla bla...
func GetAllFeedbacks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var feedbacks mymodels.AllFeedbacks
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, ActivityID, StudentID FROM Feedbacks")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, ActivityID, StudentID int
		if err := rows.Scan(&ID, &ActivityID, &StudentID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var feedback = mymodels.Feedback{ID: &ID, ActivityID: &ActivityID, StudentID: &StudentID}
		feedbacks = append(feedbacks, feedback)
	}
	json.NewEncoder(w).Encode(feedbacks)
	w.WriteHeader(http.StatusOK)
	return
}

// GetFeedback bla bla...
func GetFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	feedbackID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	var ID, ActivityID, StudentID int
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, ActivityID, StudentID FROM Feedbacks WHERE ID=?", feedbackID).Scan(&ID, &ActivityID, &StudentID)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var feedback = mymodels.Feedback{ID: &ID, ActivityID: &ActivityID, StudentID: &StudentID}
	json.NewEncoder(w).Encode(feedback)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateFeedback bla bla...
func UpdateFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedFeedback mymodels.Feedback
	json.Unmarshal(reqBody, &updatedFeedback)
	switch {
	case (updatedFeedback.ID == nil) || (*updatedFeedback.ID*1 == 0) || (*updatedFeedback.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedFeedback.ActivityID == nil) || (*updatedFeedback.ActivityID*1 == 0) || (*updatedFeedback.ActivityID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ActivityID is empty or not valid")
		return
	case (updatedFeedback.StudentID == nil) || (*updatedFeedback.StudentID*1 == 0) || (*updatedFeedback.StudentID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "StudentID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE Feedbacks SET ActivityID=?, StudentID=? WHERE ID=?", updatedFeedback.ActivityID, updatedFeedback.StudentID, updatedFeedback.ID)
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
			json.NewEncoder(w).Encode(updatedFeedback)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}

// DeleteFeedback bla bla...
func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedFeedback mymodels.Feedback
	json.Unmarshal(reqBody, &deletedFeedback)

	if (deletedFeedback.ID) == nil || (*deletedFeedback.ID*1 == 0) || (*deletedFeedback.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM Feedbacks WHERE ID=?", deletedFeedback.ID)
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
