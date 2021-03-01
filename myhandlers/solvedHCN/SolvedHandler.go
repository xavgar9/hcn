package solvedhcn

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/myhandlers/courses"
	"hcn/myhandlers/hcn"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strconv"
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

	// find all students of the course
	students, err := courses.GetAllStudentsCourseNoHTTP(*newSolvedHCN.CourseID)
	if err != nil {
		fmt.Println("2 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	students = append(students, mymodels.Student{ID: newSolvedHCN.TeacherID})
	// find the mongo id of the HCN
	hcnMongoID, err := hcn.GetHCNMongoIDNoHTTP(*newSolvedHCN.OriginalHCN)
	if err != nil {
		fmt.Println("3 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	// find the mongo document of the HCN
	hcnMongo, err := hcn.GetHCNMongoNoHTTP(hcnMongoID)
	if err != nil {
		fmt.Println("4 ", err.Error(), hcnMongoID)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	var Db, _ = config.MYSQLConnection()
	endpoint := "http://" + config.ServerIP + ":" + config.ServerPort + "/HCN/CreateHCNMongo"
	jsonValue, _ := json.Marshal(hcnMongo)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	hcnCreated := 0
	for _, student := range students {
		req, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			break
		} else {
			if req.Status == "201 Created" {
				var hcnID string
				reqBody, err := ioutil.ReadAll(req.Body)
				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "(USER) %v", err.Error())
					return
				}
				json.Unmarshal(reqBody, &hcnID)

				rows, err := Db.Exec("INSERT INTO Solved_HCN(OriginalHCN, MongoID, Solver, Reviewed) VALUES (?,?,?,?)", *newSolvedHCN.OriginalHCN, hcnID, *student.ID, 0)
				defer Db.Close()
				if err == nil {
					cnt, _ := rows.RowsAffected()
					if cnt == 1 {
						hcnCreated++
					}
				} else {
					fmt.Fprintf(w, err.Error())
					break
				}
			}
		}
		defer req.Body.Close()
	}
	if len(students) == hcnCreated {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "HCNs created: "+strconv.Itoa(hcnCreated))
	return
}

// GetAllSolvedHCN creates blank hcn for the students. Takes the
// mongoID of the original HCN and creates as copies as students in the course
func GetAllSolvedHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allSolvedHCN mymodels.AllSolvedHCN
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, OriginalHCN, MongoID, Solver, Reviewed FROM Solved_HCN")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, OriginalHCN, Solver, Reviewed int
		var MongoID string
		if err := rows.Scan(&ID, &OriginalHCN, &MongoID, &Solver, &Reviewed); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(MongoID)

		var solvedHCN = mymodels.SolvedHCN{ID: &ID, OriginalHCN: &OriginalHCN, MongoID: &MongoID, Solver: &Solver, Reviewed: &Reviewed}
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
