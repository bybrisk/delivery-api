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

func UpdatePendingDelivery (docID string, pendingCount string) int64 {

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
		log.Error("UpdatePendingDelivery ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func DecreasePendingDelivery (docID string, pendingCount string) int64 {

	pendingInt, err := strconv.Atoi(pendingCount)
    if err != nil {
        log.Error("updatePendingDelivery str to int ERROR:")
		log.Error(err)
	}

	newPendingInt := pendingInt - 1
	newPendingstring := strconv.Itoa(newPendingInt)
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"deliveryPending": newPendingstring}})
	if err != nil {
		log.Error("DecreasePendingDelivery ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func UpdateTransitDelivery (docID string, transitCount string) int64 {

	transitInt, err := strconv.Atoi(transitCount)
    if err != nil {
        log.Error("UpdateTransitDelivery str to int ERROR:")
		log.Error(err)
	}

	newTransitInt := transitInt + 1
	newTransitstring := strconv.Itoa(newTransitInt)
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"deliveryTransit": newTransitstring}})
	if err != nil {
		log.Error("UpdateTransitDelivery ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func DecreaseTransitDelivery (docID string, transitCount string) int64 {

	transitInt, err := strconv.Atoi(transitCount)
    if err != nil {
        log.Error("DecreaseTransitDelivery str to int ERROR:")
		log.Error(err)
	}

	newTransitInt := transitInt - 1
	newTransitstring := strconv.Itoa(newTransitInt)
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"deliveryTransit": newTransitstring}})
	if err != nil {
		log.Error("DecreaseTransitDelivery ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func UpdateCancelledDelivery (docID string, cancelledCount string) int64 {

	cancelledInt, err := strconv.Atoi(cancelledCount)
    if err != nil {
        log.Error("UpdateCancelledDelivery str to int ERROR:")
		log.Error(err)
	}

	newCancelledInt := cancelledInt + 1
	newCancelledstring := strconv.Itoa(newCancelledInt)
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"deliveryCancelled": newCancelledstring}})
	if err != nil {
		log.Error("UpdateCancelledDelivery ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func UpdateDeliveredDelivery (docID string, deleveredCount string) int64 {

	deleveredInt, err := strconv.Atoi(deleveredCount)
    if err != nil {
        log.Error("UpdateDeliveredDelivery str to int ERROR:")
		log.Error(err)
	}

	newDeleveredInt := deleveredInt + 1
	newDeleveredstring := strconv.Itoa(newDeleveredInt)
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"deliveryDelivered": newDeleveredstring}})
	if err != nil {
		log.Error("UpdateDeliveredDelivery ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func GetDeliveryFrequency (docID string) *DeliveryCountStatus {
	var count *DeliveryCountStatus
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	
	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&count)
	if err != nil {
		log.Error("GetDeliveryFrequency ERROR:")
		log.Error(err)
	}
	return count
}