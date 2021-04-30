package mongoHelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CleanHCN Create the same struct of the HCN without
// the mongo _id
func CleanHCN(hcn mymodels.HCNmongo) mymodels.HCNmongoNoID {
	var newHCN mymodels.HCNmongoNoID
	hcnJSON, _ := json.Marshal(hcn)
	json.Unmarshal([]byte(hcnJSON), &newHCN)
	return newHCN
}

// CreateHCNMongo creates an HCN in Mongo db
func CreateHCNMongo(hcnMongo mymodels.HCNmongo) (string, error) {
	var mongoID string
	// Insert data in mongo db
	client, ctx := config.MongoConnection()
	collection := client.Database(config.MongoDB).Collection(config.MongoCollection)
	result, err := collection.InsertOne(ctx, CleanHCN(hcnMongo))
	if err != nil {
		return mongoID, err
	}
	mongoID = strings.Split(fmt.Sprintf("%v", result.InsertedID), `"`)[1]
	return mongoID, nil
}

// GetHCN returns the mongo document that corresponds to the _id
func GetHCN(id string) (mymodels.HCNmongo, error) {
	var newHCN mymodels.HCNmongo

	//hcnID, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("Noooio 1")
	hcnID, err := primitive.ObjectIDFromHex(id)
	fmt.Println("Noooio 2")
	if err != nil {
		return newHCN, errors.New("Id is empty or not valid")
	}
	fmt.Println("Noooio 3")
	client, ctx := config.MongoConnection()
	collection := client.Database(config.MongoDB).Collection(config.MongoCollection)
	fmt.Println("Noooio 3")
	collection.FindOne(ctx, bson.M{"_id": hcnID}).Decode(&newHCN)
	fmt.Println("Noooio 4")
	fmt.Println(newHCN)
	if newHCN.ID == nil {
		return newHCN, errors.New("Mongo HCN doesnt exist")
	}
	return newHCN, nil
}
