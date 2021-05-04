package solvedhcn

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/helpers/hcnHelper"
	"hcn/helpers/mongoHelper"
	"hcn/myhandlers/courses"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strconv"

	dbHelper "hcn/myhelpers/databaseHelper"
)

// CreateSolvedHCN creates blank hcn for the students. Takes the
// mongoID of the original HCN and creates as copies as students in the course
func CreateSolvedHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newSolvedHCN mymodels.SolvedHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("1 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	json.Unmarshal(reqBody, &newSolvedHCN)
	switch {
	case (newSolvedHCN.CourseID == nil) || (*newSolvedHCN.CourseID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "CourseID is empty or not valid")
		return
	case (newSolvedHCN.ActivityID == nil) || (*newSolvedHCN.ActivityID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ActivityID is empty or not valid")
		return
	case (newSolvedHCN.OriginalHCN == nil) || (*newSolvedHCN.OriginalHCN*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "OriginalHCN is empty or not valid")
		return
	case (newSolvedHCN.TeacherID == nil) || (*newSolvedHCN.TeacherID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	default:
		// find all students of the course
		students, err := courses.GetAllStudentsCourseNoHTTP(*newSolvedHCN.CourseID)
		if err != nil {
			fmt.Println("2 ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		// the teacher could solve the activity
		students = append(students, mymodels.Student{ID: newSolvedHCN.TeacherID})

		// find the mongo id of the HCN
		mongoID, err := hcnHelper.GetSqlID(*newSolvedHCN.OriginalHCN)
		//hcnMongoID, err := hcn.GetHCNMongoIDNoHTTP(*newSolvedHCN.OriginalHCN)
		if err != nil {
			fmt.Println("3 ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		// find the mongo document of the HCN
		hcnMongo, err := mongoHelper.GetHCN(mongoID)
		if err != nil {
			fmt.Println("4 ", err.Error(), mongoID)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, err.Error())
			return
		}

		hcnCreated := 0
		for _, student := range students {
			mongoID, err := mongoHelper.CreateHCNMongo(hcnMongo)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("5 ", err.Error())
				fmt.Fprintf(w, err.Error())
				return
			} else {
				Db, err := config.MYSQLConnection()
				defer Db.Close()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, err.Error())
					return
				}
				rows, err := Db.Exec("INSERT INTO Solved_HCN(ActivityID, OriginalHCN, MongoID, Solver, Reviewed) VALUES (?,?,?,?,?)", *newSolvedHCN.ActivityID, *newSolvedHCN.OriginalHCN, mongoID, *student.ID, 0)
				if err == nil {
					cnt, _ := rows.RowsAffected()
					if cnt == 1 {
						hcnCreated++
					}
				}
			}
		}
		if len(students) == hcnCreated {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "HCNs created: "+strconv.Itoa(hcnCreated))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}

// GetAllSolvedHCN creates blank hcn for the students. Takes the
// mongoID of the original HCN and creates as copies as students in the course
func GetAllSolvedHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	activityID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	fmt.Println("ActivityID", activityID)
	var allSolvedHCN mymodels.AllSolvedHCN
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, ActivityID, OriginalHCN, MongoID, Solver, Reviewed FROM Solved_HCN WHERE ActivityID=?", activityID)
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var id, activityID, originalHCN, solver, reviewed int
		var mongoID string
		if err := rows.Scan(&id, &activityID, &originalHCN, &mongoID, &solver, &reviewed); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(mongoID)

		var solvedHCN = mymodels.SolvedHCN{ID: &id, ActivityID: &activityID, OriginalHCN: &originalHCN, MongoID: &mongoID, Solver: &solver, Reviewed: &reviewed}
		allSolvedHCN = append(allSolvedHCN, solvedHCN)
	}
	if allSolvedHCN == nil {
		var emptyTest mymodels.EmptyTest
		json.NewEncoder(w).Encode(emptyTest)
	} else {
		// Use next line for testing
		//json.NewEncoder(w).Encode(allSolvedHCN)
		// Use next line for production
		json.NewEncoder(w).Encode(allSolvedHCN)
	}
	return
}

// UpdateSolvedHCN updates the reviewed status of an hcn,
func UpdateSolvedHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedSolvedHCN mymodels.SolvedHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedSolvedHCN)

	// Fields validation
	structFields := []string{"ActivityID", "Solver", "Reviewed"} // struct fields to check
	_, err = updatedSolvedHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedSolvedHCN)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, err.Error())
	}

	return
}

// DeleteSolvedHCN all solved hcns filtered by activity id.
func DeleteSolvedHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedSolvedHCN mymodels.SolvedHCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedSolvedHCN)

	// Fields validation
	structFields := []string{"ActivityID"} // struct fields to check
	_, err = deletedSolvedHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedSolvedHCN)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}

	return
}
