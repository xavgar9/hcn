package hcn

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"strings"

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
	case (newHCN.TeacherID == nil) || (*newHCN.TeacherID*1 <= 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	case (newHCN.MongoID == nil) || (len(*newHCN.MongoID) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "MongoID is empty or not valid")
		return
	default:
		rows, err := Db.Exec("INSERT INTO HCN(TeacherID, MongoID) VALUES (?, ?)", newHCN.TeacherID, newHCN.MongoID)
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
	rows, err := Db.Query("SELECT ID, TeacherID, MongoID FROM HCN")
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
		var MongoID string
		if err := rows.Scan(&ID, &TeacherID, &MongoID); err != nil {
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var hcn = mymodels.HCN{ID: &ID, TeacherID: &TeacherID, MongoID: &MongoID}
		hcns = append(hcns, hcn)
	}
	json.NewEncoder(w).Encode(hcns)
	w.WriteHeader(http.StatusOK)
	return
}

// GetHCNMongoIDNoHTTP MYSQL bla bla...
func GetHCNMongoIDNoHTTP(hcnID int) (string, error) {
	var MongoID string
	var Db, _ = config.MYSQLConnection()
	err := Db.QueryRow("SELECT MongoID FROM HCN WHERE ID=?", hcnID).Scan(&MongoID)
	defer Db.Close()
	if err != nil {
		return MongoID, err
	}
	return MongoID, err
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
	var MongoID string
	var Db, _ = config.MYSQLConnection()
	err = Db.QueryRow("SELECT ID, TeacherID, MongoID FROM HCN WHERE ID=?", hcnID).Scan(&ID, &TeacherID, &MongoID)
	defer Db.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var course = mymodels.HCN{ID: &ID, TeacherID: &TeacherID, MongoID: &MongoID}
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
	case (updatedHCN.MongoID == nil) || (len(*updatedHCN.MongoID) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "MongoID is empty or not valid")
		return
	default:
		var Db, _ = config.MYSQLConnection()
		row, err := Db.Exec("UPDATE HCN SET TeacherID=?, MongoID=? WHERE ID=?", updatedHCN.TeacherID, updatedHCN.MongoID, updatedHCN.ID)
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

// CreateHCNMongo Mongo bla bla... OK
func CreateHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}

	// Prepare the data for insert
	reqEmpty := true
	var newGeneralData *mymodels.GeneralData
	var newReason *string
	var newPatientData *mymodels.PatientData
	var newAnthropometry *mymodels.Anthropometry
	var newBiochemistry *[]mymodels.Biochemistry

	if gjson.Get(string(reqBody), "GeneralData").Exists() {
		generalData := gjson.Get(string(reqBody), "GeneralData")
		json.Unmarshal([]byte(generalData.Raw), &newGeneralData)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "ConsultationReason").Exists() {
		consultationReason := gjson.Get(string(reqBody), "ConsultationReason")
		json.Unmarshal([]byte(consultationReason.Raw), &newReason)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "Anthropometry").Exists() {
		anthropometry := gjson.Get(string(reqBody), "Anthropometry")
		json.Unmarshal([]byte(anthropometry.Raw), &newAnthropometry)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "Biochemistry").Exists() {
		biochemistry := gjson.Get(string(reqBody), "Biochemistry")
		json.Unmarshal([]byte(biochemistry.Raw), &newBiochemistry)
		reqEmpty = false
	}

	// This strange way of initialize the struct if looking for
	// no storing empty fields in mongo
	newHCNmongo := mymodels.HCNmongo{
		GeneralData:        *&newGeneralData,
		ConsultationReason: *&newReason,
		PatientData:        *&newPatientData,
		Anthropometry:      *&newAnthropometry,
		Biochemistry:       *&newBiochemistry,
	}

	if !reqEmpty {
		// Insert data in mongo db
		client, ctx := config.MongoConnection()
		collection := client.Database("HCNProject").Collection("HCN")
		result, err := collection.InsertOne(ctx, newHCNmongo)
		if err != nil {
			w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		}
		/*
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
		*/

		w.WriteHeader(http.StatusCreated)

		// Use next line for testing
		//json.NewEncoder(w).Encode(newHCNmongo)
		// Use next line for production
		json.NewEncoder(w).Encode(strings.Split(fmt.Sprintf("%v", result.InsertedID), `"`)[1])
	}

}

// GetAllHCNMongo bla bla... OK
func GetAllHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Use next line for testing
	//var allHCNsNoID mymodels.AllHCNmongoNoID

	// Use next line for production
	var allHCNs mymodels.AllHCNmongo

	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "(SQL) %v", err.Error())
		w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var newHCN mymodels.HCNmongo
		cursor.Decode(&newHCN)

		// Use next line for testing
		// allHCNsNoID = append(allHCNsNoID, helpers.CleanHCN(newHCN))

		// Use next line for production
		allHCNs = append(allHCNs, newHCN)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}
	// Use next line for testing
	//if allHCNsNoID == nil {
	// Use next line for production
	if allHCNs == nil {
		var emptyTest mymodels.EmptyTest
		json.NewEncoder(w).Encode(emptyTest)
	} else {
		// Use next line for testing
		//json.NewEncoder(w).Encode(allHCNsNoID)
		// Use next line for production
		json.NewEncoder(w).Encode(allHCNs)
	}

}

// GetHCNMongoNoHTTP bla bla... OK
func GetHCNMongoNoHTTP(hcnid string) (mymodels.HCNmongo, error) {
	hcnID, _ := primitive.ObjectIDFromHex(hcnid)

	var newHCN mymodels.HCNmongo

	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")
	collection.FindOne(ctx, bson.M{"_id": hcnID}).Decode(&newHCN)
	if newHCN.ID == nil {
		return newHCN, errors.New("Mongo HCN doesnt exist")
	}
	return newHCN, nil

}

// GetHCNMongo bla bla... OK
func GetHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	hcnID, _ := primitive.ObjectIDFromHex(keys[0])

	var newHCN mymodels.HCNmongo

	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")
	collection.FindOne(ctx, bson.M{"_id": hcnID}).Decode(&newHCN)
	if newHCN.ID == nil {
		json.NewEncoder(w).Encode(nil)
	} else {
		// For testing only
		var newHCNnoID mymodels.HCNmongoNoID
		newhcn, _ := json.Marshal(newHCN)
		json.Unmarshal([]byte(newhcn), &newHCNnoID)
		json.NewEncoder(w).Encode(newHCNnoID)

		//Use this in production
		//json.NewEncoder(w).Encode(newHCN)
	}

}

// UpdateHCNMongo ...
// It is neccesary that all the fields inside a field are filled, otherwise
// the unfilled fields inside a field will be erased.
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

	// All json assessment data
	json.Unmarshal(reqBody, &newHCN)

	// General nutritional assessment data
	generalData := gjson.Get(string(reqBody), "GeneralData")
	json.Unmarshal([]byte(generalData.Raw), &newGeneralData)

	// Patient data
	patientData := gjson.Get(string(reqBody), "PatientData")
	json.Unmarshal([]byte(patientData.Raw), &newPatientData)

	// Anthropometry data
	Anthropometry := gjson.Get(string(reqBody), "Anthropometry")
	json.Unmarshal([]byte(Anthropometry.Raw), &newAnthropometry)

	// Biochemistry data
	biochemistry := gjson.Get(string(reqBody), "Biochemistry")
	json.Unmarshal([]byte(biochemistry.Raw), &newBiochemistry)

	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")

	id, _ := primitive.ObjectIDFromHex(*newHCN.ID)
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{
				{"GeneralData", newGeneralData},
				{"PatientData", newPatientData},
				{"Anthropometry", newAnthropometry},
				{"Biochemistry", newBiochemistry}}},
		},
	)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
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
		id, _ := primitive.ObjectIDFromHex(*newHCN.ID)
		_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
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
	//json.NewEncoder(w).Encode(cnt)
}

////////////////////////////////////////
// HAY QUE METERLE SQL A TODO MONGO HCN
////////////////////////////////////////

// DeleteHCNMongo bla bla...
func DeleteHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	type deletedID struct {
		ID *string
	}
	var deletedHCN deletedID
	json.Unmarshal(reqBody, &deletedHCN)

	if (deletedHCN.ID == nil) || (len(*deletedHCN.ID) == 0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	client, ctx := config.MongoConnection()
	collection := client.Database("HCNProject").Collection("HCN")
	id, _ := primitive.ObjectIDFromHex(*deletedHCN.ID)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.DeletedCount)
}
