package data

import (
	"fmt"
	"github.com/bybrisk/structs"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	log "github.com/sirupsen/logrus"
)

var resultID string

//Database Funcs
func createBusinessAccount (account *BusinessAccountRequest) string {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	result, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, account)
	if insertErr != nil {
		log.Error("Create Business Account ERROR:")
		log.Error(insertErr)
	} else {
		fmt.Println("createBusinessAccount() API result:", result)

		newID := result.InsertedID
		fmt.Println("createBusinessAccount() newID:", newID)
		resultID = newID.(primitive.ObjectID).Hex()
	}
	return resultID
}

func getProfileConfig (account *BusinessAccountRequest) *structs.ProfileConfig {
	collectionName := shashankMongo.DatabaseName.Collection("profileConfig")
	filter := bson.M{"plan": account.BusinessPlan}
	var document *structs.ProfileConfig

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("setProfileConfig ERROR:")
		log.Error(err)
	}
	return document
}	

func setProfileConfig (document *structs.ProfileConfig, docID string) int64 {
	//update businessAccount
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	update := bson.M{"$set": bson.M{"profileConfig": document}}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("UpdateDeliveryInfo ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount
}

func setDeliveryStats (docID string) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	update := bson.M{"$set":bson.M{"deliveryPending": "0", "deliveryDelivered":"0"}}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("setDeliveryStats ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount
}

func getBusinessAccount (docID string) *BusinessAccountResponse {
	var businessAccount *BusinessAccountResponse
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
    err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&businessAccount)
	if err != nil {
		log.Error("getBusinessAccount ERROR:")
		log.Error(err)
	}

	return businessAccount
}

func updateBusinessAccount(account *UpdateBusinessAccountRequest) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(account.BybID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.ReplaceOne(shashankMongo.CtxForDB, filter, account)
	if err != nil {
		log.Error("updateBusinessAccount ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}
