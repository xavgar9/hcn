package hcn

import (
	"encoding/json"
	"fmt"
	"hcn/config"
	hcnHelper "hcn/helpers/hcnHelper"
	mongoHelper "hcn/helpers/mongoHelper"
	"hcn/mymodels"
	"strings"

	"io/ioutil"
	"net/http"

	dbHelper "hcn/myhelpers/databaseHelper"
	stuctHelper "hcn/myhelpers/structValidationHelper"

	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateHCN MySQL creates one HCN in db.
func CreateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newHCN mymodels.HCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &newHCN)

	// Fields validation
	structFields := []string{"TeacherID", "MongoID"} // struct fields to check
	_, err = newHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Insert(newHCN)
	if err != nil {
		if strings.Split(err.Error(), ":")[0] == "(db 2) Error 1062" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

// GetAllHCN MySQL bla bla...
func GetAllHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Data from db
	var hcn mymodels.HCN
	rows, err := dbHelper.GetAll(hcn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var allHCN mymodels.AllHCN
	dbHelper.RowsToStruct(rows, &allHCN)

	json.NewEncoder(w).Encode(allHCN)
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

// GetHCN MySQL returns one hcn filtered by the id.
func GetHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := stuctHelper.GetURLParameter("ID", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	hcn := mymodels.HCN{ID: &id}

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = hcn.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data from into db
	rows, err := dbHelper.Get(hcn)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
		return
	}
	var allHCN mymodels.AllHCN
	dbHelper.RowsToStruct(rows, &allHCN)
	json.NewEncoder(w).Encode(allHCN[0])

	return
}

// UpdateHCN MySQL updates fields of an HCN in db.
func UpdateHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedHCN mymodels.HCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedHCN)

	// Fields validation
	_, err = updatedHCN.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data update into db
	_, err = dbHelper.Update(updatedHCN)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
	}
	return
}

// DeleteHCN MySQL deletes one HCN filtered by the id.
func DeleteHCN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var deletedHCN mymodels.HCN
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	json.Unmarshal(reqBody, &deletedHCN)

	// Fields validation
	structFields := []string{"ID"} // struct fields to check
	_, err = deletedHCN.ValidateFields(structFields)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Data insertion into db
	_, err = dbHelper.Delete(deletedHCN)
	if err != nil {
		if string(err.Error()[4]) == "2" {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, err.Error())
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
	var newConsultationReason *string
	var newPatientData *mymodels.PatientData
	var newAnthropometry *mymodels.Anthropometry
	var newGeneralInterpretation *string
	var newGeneralFeedback *string
	var teacherID *int

	if gjson.Get(string(reqBody), "GeneralData").Exists() {
		generalData := gjson.Get(string(reqBody), "GeneralData")
		json.Unmarshal([]byte(generalData.Raw), &newGeneralData)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "PatientData").Exists() {
		patientData := gjson.Get(string(reqBody), "PatientData")
		json.Unmarshal([]byte(patientData.Raw), &newPatientData)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "ConsultationReason").Exists() {
		consultationReason := gjson.Get(string(reqBody), "ConsultationReason")
		json.Unmarshal([]byte(consultationReason.Raw), &newConsultationReason)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "Anthropometry").Exists() {
		anthropometry := gjson.Get(string(reqBody), "Anthropometry")
		json.Unmarshal([]byte(anthropometry.Raw), &newAnthropometry)
		reqEmpty = false
	}

	var newAllBiochemistryParameters mymodels.AllBiochemistryParameters
	var interpretation *string
	var feedback *string
	if gjson.Get(string(reqBody), "Biochemistry").Exists() {

		// Look for the array of parameters
		biochemistryParameters := gjson.Get(string(reqBody), "Biochemistry.Parameters")
		// Loop the array of parameters
		for _, parameter := range biochemistryParameters.Array() {
			var newParameter mymodels.BiochemistryParameters
			json.Unmarshal([]byte(parameter.Raw), &newParameter)
			newAllBiochemistryParameters = append(newAllBiochemistryParameters, newParameter)
		}

		biochemistryInterpretation := gjson.Get(string(reqBody), "Biochemistry.Interpretation")
		biochemistryFeedback := gjson.Get(string(reqBody), "Biochemistry.Feedback")
		json.Unmarshal([]byte(biochemistryInterpretation.Raw), &interpretation)
		json.Unmarshal([]byte(biochemistryFeedback.Raw), &feedback)
		reqEmpty = false
	}

	if gjson.Get(string(reqBody), "Interpretation").Exists() {
		interpretation := gjson.Get(string(reqBody), "Interpretation")
		json.Unmarshal([]byte(interpretation.Raw), &newGeneralInterpretation)
	}

	if gjson.Get(string(reqBody), "Feedback").Exists() {
		feedback := gjson.Get(string(reqBody), "Feedback")
		json.Unmarshal([]byte(feedback.Raw), &newGeneralFeedback)
	}

	if gjson.Get(string(reqBody), "TeacherID").Exists() {
		teacherid := gjson.Get(string(reqBody), "TeacherID")
		json.Unmarshal([]byte(teacherid.Raw), &teacherID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "TeacherID is empty or not valid")
		return
	}

	// We have to create the biochemistry struct like this
	// because it has an array of structs
	newBiochemistry := mymodels.Biochemistry{
		Parameters:     &newAllBiochemistryParameters,
		Interpretation: interpretation,
		Feedback:       feedback,
	}

	// This strange way of initialize the struct is looking for
	// no storing empty fields in mongo
	newHCNmongo := mymodels.HCNmongo{
		GeneralData:        *&newGeneralData,
		ConsultationReason: *&newConsultationReason,
		PatientData:        *&newPatientData,
		Anthropometry:      *&newAnthropometry,
		Biochemistry:       &newBiochemistry,
		Interpretation:     *&newGeneralInterpretation,
		Feedback:           *&newGeneralFeedback,
	}

	if !reqEmpty {

		// Insert data in mongo db
		mongoID, err := mongoHelper.CreateHCNMongo(newHCNmongo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		fmt.Println("Datos antes", *teacherID, mongoID)

		// Insert data in mongo db
		hcn := mymodels.HCN{TeacherID: teacherID, MongoID: &mongoID}
		_, err = hcnHelper.CreateHCN(hcn)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "(SQL) %v", err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		// Use next line for testing
		//json.NewEncoder(w).Encode(newHCNmongo)
		// Use next line for production
		json.NewEncoder(w).Encode(mongoID)
		return

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
	collection := client.Database(config.MongoDB).Collection(config.MongoCollection)

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
	// Use next "if" line for testing
	//if allHCNsNoID == nil {
	// Use next "if" line for production
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

// GetHCNMongo bla bla... OK
func GetHCNMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("1")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is empty or not valid")
		return
	}
	fmt.Println("2")
	newHCN, err := mongoHelper.GetHCN(keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	// For testing only
	// var newHCNnoID mymodels.HCNmongoNoID
	// newhcn, _ := json.Marshal(newHCN)
	// json.Unmarshal([]byte(newhcn), &newHCNnoID)
	// json.NewEncoder(w).Encode(newHCNnoID)

	//Use this in production
	json.NewEncoder(w).Encode(newHCN)
	return
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
	var newGeneralData *mymodels.GeneralData
	var newPatientData *mymodels.PatientData
	var newConsultationReason *string
	var newAnthropometry *mymodels.Anthropometry

	var newGeneralInterpretation *string
	var newGeneralFeedback *string

	// All json assessment data
	json.Unmarshal(reqBody, &newHCN)

	// General nutritional assessment data
	generalData := gjson.Get(string(reqBody), "GeneralData")
	json.Unmarshal([]byte(generalData.Raw), &newGeneralData)

	// Patient data
	patientData := gjson.Get(string(reqBody), "PatientData")
	json.Unmarshal([]byte(patientData.Raw), &newPatientData)

	// Consultation reason nutritional assessment data
	consultationReason := gjson.Get(string(reqBody), "ConsultationReason")
	json.Unmarshal([]byte(consultationReason.Raw), &newConsultationReason)

	// Anthropometry data
	Anthropometry := gjson.Get(string(reqBody), "Anthropometry")
	json.Unmarshal([]byte(Anthropometry.Raw), &newAnthropometry)

	// Biochemistry data
	var newAllBiochemistryParameters mymodels.AllBiochemistryParameters
	var interpretation *string
	var feedback *string
	// Look for the array of parameters
	biochemistryParameters := gjson.Get(string(reqBody), "Biochemistry.Parameters")
	// Loop the array of parameters
	for _, parameter := range biochemistryParameters.Array() {
		var newParameter mymodels.BiochemistryParameters
		json.Unmarshal([]byte(parameter.Raw), &newParameter)
		newAllBiochemistryParameters = append(newAllBiochemistryParameters, newParameter)
	}

	biochemistryInterpretation := gjson.Get(string(reqBody), "Biochemistry.Interpretation")
	biochemistryFeedback := gjson.Get(string(reqBody), "Biochemistry.Feedback")
	json.Unmarshal([]byte(biochemistryInterpretation.Raw), &interpretation)
	json.Unmarshal([]byte(biochemistryFeedback.Raw), &feedback)

	// We have to create the biochemistry struct like this
	// because it has an array of structs
	newBiochemistry := mymodels.Biochemistry{
		Parameters:     &newAllBiochemistryParameters,
		Interpretation: interpretation,
		Feedback:       feedback,
	}

	// General interpretation data
	generalInterpretation := gjson.Get(string(reqBody), "Interpretation")
	json.Unmarshal([]byte(generalInterpretation.Raw), &newGeneralInterpretation)

	// General feedback data
	generalFeedback := gjson.Get(string(reqBody), "Feedback")
	json.Unmarshal([]byte(generalFeedback.Raw), &newGeneralFeedback)

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
				{"ConsultationReason", newConsultationReason},
				{"Anthropometry", newAnthropometry},
				{"Biochemistry", newBiochemistry},
				{"Interpretation", newGeneralInterpretation},
				{"Feedback", newGeneralFeedback}}},
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
