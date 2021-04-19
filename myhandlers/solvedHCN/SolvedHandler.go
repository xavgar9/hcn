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
	fmt.Println("Solved 1")
	students, err := courses.GetAllStudentsCourseNoHTTP(*newSolvedHCN.CourseID)
	if err != nil {
		fmt.Println("2 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	students = append(students, mymodels.Student{ID: newSolvedHCN.TeacherID})

	// find the mongo id of the HCN
	fmt.Println("Solved 2")
	hcnMongoID, err := hcn.GetHCNMongoIDNoHTTP(*newSolvedHCN.OriginalHCN)
	if err != nil {
		fmt.Println("3 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	// find the mongo document of the HCN
	fmt.Println("Solved 3")
	hcnMongo, err := hcn.GetHCNMongoNoHTTP(hcnMongoID)
	if err != nil {
		fmt.Println("4 ", err.Error(), hcnMongoID)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	endpoint := "http://" + config.ServerIP + ":" + config.ServerPort + "/HCN/CreateHCNMongo"
	jsonValue, err := json.Marshal(hcnMongo)
	fmt.Println("Solved 4")
	if err != nil {
		fmt.Println("5 ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Println("Solved 5")
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("6 ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	fmt.Println("Solved 6")
	hcnCreated := 0
	for _, student := range students {
		fmt.Println("Solved 7")
		req, err := client.Do(req)
		fmt.Println("Solved 7.1", *student.ID)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			break
		} else {
			if req.Status == "201 Created" {
				var hcnID string
				reqBody, err := ioutil.ReadAll(req.Body)
				if err != nil {
					fmt.Println("Error here ", err.Error())
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "(USER) %v", err.Error())
					return
				}
				json.Unmarshal(reqBody, &hcnID)
				var Db, _ = config.MYSQLConnection()
				defer Db.Close()
				rows, err := Db.Exec("INSERT INTO Solved_HCN(OriginalHCN, MongoID, Solver, Reviewed) VALUES (?,?,?,?)", *newSolvedHCN.OriginalHCN, hcnID, *student.ID, 0)
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
