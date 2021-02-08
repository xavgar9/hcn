package hcn

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateHCN MySQL bla bla...
func CreateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newHCN mymodels.HCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	var Db, _ = config.MYSQLConnection()
	json.Unmarshal(reqBody, &newHCN)
	switch {
	case (newHCN.TeacherID == nil) || (*newHCN.TeacherID*1 == 0) || (*newHCN.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO HCN(TeacherID) VALUES (?)", newHCN.TeacherID)
		defer Db.Close()
		if err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		cnt, _ := rows.RowsAffected()
		if cnt == 1 {
			int64ID, _ := rows.LastInsertId()
			intID := int(int64ID)
			newHCN.ID = &intID
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newHCN)
		}
		return
	}
}

// GetAllHCN MySQL bla bla...
func GetAllHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hcns mymodels.AllHCN
	var Db, _ = config.MYSQLConnection()
	rows, err := Db.Query("SELECT ID, TeacherID FROM HCN")
	defer Db.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for rows.Next() {
		var ID, TeacherID int
		if err := rows.Scan(&ID, &TeacherID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var hcn = mymodels.HCN{ID: &ID, TeacherID: &TeacherID}
		hcns = append(hcns, hcn)
	}
	json.NewEncoder(w).Encode(hcns)
	w.WriteHeader(http.StatusOK)
	return
}

// GetHCN MySQL bla bla...
func GetHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	hcnID, err := strconv.Atoi(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	var ID, TeacherID int
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, TeacherID FROM HCN WHERE ID=?", hcnID).Scan(&ID, &TeacherID)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var course = mymodels.HCN{ID: &ID, TeacherID: &TeacherID}
	json.NewEncoder(w).Encode(course)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateHCN MySQL bla bla...
func UpdateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var updatedHCN mymodels.HCN
	json.Unmarshal(reqBody, &updatedHCN)
	switch {
	case (updatedHCN.ID == nil) || (*updatedHCN.ID*1 == 0) || (*updatedHCN.ID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	case (updatedHCN.TeacherID == nil) || (*updatedHCN.TeacherID*1 == 0) || (*updatedHCN.TeacherID*1 < 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE HCN SET TeacherID=? WHERE ID=?", updatedHCN.TeacherID, updatedHCN.ID)
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
			json.NewEncoder(w).Encode(updatedHCN)
		} else {
			fmt.Fprintf(w, "No rows updated")
		}
		return
	}
}

// DeleteHCN MySQL bla bla...
func DeleteHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	var deletedHCN mymodels.HCN
	json.Unmarshal(reqBody, &deletedHCN)

	if (deletedHCN.ID) == nil || (*deletedHCN.ID*1 == 0) || (*deletedHCN.ID*1 < 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}

	var Db, _ = config.MYSQLConnection()
	row, err := Db.Exec("DELETE FROM HCN WHERE ID=?", deletedHCN.ID)
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

// -------------------------------------------
// Mongo Handlers
// -------------------------------------------

// CreateHCNMongo Mongo bla bla...
func CreateHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	// Prepare the data for insert
	var newHCNmongo mymodels.HCNmongo
	var newGeneralData mymodels.GeneralData
	var newPatientData mymodels.PatientData
	var newAnthropometry mymodels.Anthropometry
	var newBiochemistry []mymodels.Biochemistry

	generalData := gjson.Get(string(reqBody), "GeneralData")
	consultationReason := gjson.Get(string(reqBody), "ConsultationReason")
	anthropometry := gjson.Get(string(reqBody), "Anthropometry")
	patientData := gjson.Get(string(reqBody), "PatientData")
	biochemistry := gjson.Get(string(reqBody), "Biochemistry")

	// General nutritional assessment data
	json.Unmarshal([]byte(generalData.Raw), &newGeneralData)
	newHCNmongo.GeneralData = newGeneralData
	// General patient data
	json.Unmarshal([]byte(patientData.Raw), &newPatientData)
	newHCNmongo.PatientData = newPatientData
	// Consultation Reason
	reason := consultationReason.String()
	newHCNmongo.ConsultationReason = &reason
	// Anthropometry
	json.Unmarshal([]byte(anthropometry.Raw), &newAnthropometry)
	newHCNmongo.Anthropometry = newAnthropometry
	// Biochemistry
	json.Unmarshal([]byte(biochemistry.Raw), &newBiochemistry)
	newHCNmongo.Biochemistry = newBiochemistry

	// Insert data in mongo db
	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")
	result, _ := collection.InsertOne(ctx, newHCNmongo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// GetAllHCNMongo bla bla...
func GetAllHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var allHCNs mymodels.AllHCNmongo
	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		//w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var newHCN mymodels.HCNmongo
		cursor.Decode(&newHCN)
		allHCNs = append(allHCNs, newHCN)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(allHCNs)
}

// GetHCNMongo bla bla...
func GetHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	hcnID, _ := primitive.ObjectIDFromHex(keys[0])

	//var person Person
	var newHCN mymodels.HCNmongo
	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")

	err := collection.FindOne(ctx, mymodels.HCNmongo{ID: hcnID}).Decode(&newHCN)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(newHCN)
}

// UpdateHCNMongo Mongo bla bla...
func UpdateHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	// Prepare the data for insert
	var newHCN mymodels.HCNmongo
	var newGeneralData mymodels.GeneralData
	var newPatientData mymodels.PatientData
	var newAnthropometry mymodels.Anthropometry
	var newBiochemistry []mymodels.Biochemistry

	generalData := gjson.Get(string(reqBody), "GeneralData")
	consultationReason := gjson.Get(string(reqBody), "ConsultationReason")
	anthropometry := gjson.Get(string(reqBody), "Anthropometry")
	patientData := gjson.Get(string(reqBody), "PatientData")
	biochemistry := gjson.Get(string(reqBody), "Biochemistry")

	// General nutritional assessment data
	json.Unmarshal([]byte(generalData.Raw), &newGeneralData)
	newHCN.GeneralData = newGeneralData
	// General patient data
	json.Unmarshal([]byte(patientData.Raw), &newPatientData)
	newHCN.PatientData = newPatientData
	// Consultation Reason
	reason := consultationReason.String()
	newHCN.ConsultationReason = &reason
	// Anthropometry
	json.Unmarshal([]byte(anthropometry.Raw), &newAnthropometry)
	newHCN.Anthropometry = newAnthropometry
	// Biochemistry
	json.Unmarshal([]byte(biochemistry.Raw), &newBiochemistry)
	newHCN.Biochemistry = newBiochemistry

	// Update data in mongo db
	/*
		client, ctx := config.MongoConnection()
		collection := client.Database("HCNProject").Collection("HCN")
		filter := bson.M{"_id": "602078f08d054d95b0d74048"}
		update := bson.M{
			"$set": {
				"GeneralData": {
					"ValorationDate": "Malo",
				},
			},
		}

		result, _ := collection.UpdateOne(ctx, filter, update)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	*/
}

// DeleteAllHCNMongo bla bla...
func DeleteAllHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	cnt := 0
	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		//w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var newHCN mymodels.HCNmongo
		cursor.Decode(&newHCN)
		_, err := collection.DeleteOne(ctx, bson.M{"_id": newHCN.ID})
		if err != nil {
			return
		}
		cnt++
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cnt)
}
