package data

import (
	"strconv"
	//"github.com/bybrisk/structs"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	log "github.com/sirupsen/logrus"
)

var resultID string

func UpdatePendingDlivery (docID string, pendingCount string) int64 {

	pendingInt, err := strconv.Atoi(pendingCount)
    if err != nil {
        log.Error("updatePendingDelivery str to int ERROR:")
		log.Error(err)
	}

	newPendingInt := pendingInt + 1
	newPendingstring := strconv.Itoa(newPendingInt)
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"deliveryPending": newPendingstring}})
	if err != nil {
		log.Error("updateAgentsByID ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func GetPendingDelivery (docID string) *DeliveryCountStatus {
	var count *DeliveryCountStatus
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	
	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&count)
	if err != nil {
		log.Error("GetPendingDelivery ERROR:")
		log.Error(err)
	}
	return count
}