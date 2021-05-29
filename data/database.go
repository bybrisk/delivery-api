package data

import (
	"strconv"
	//"github.com/bybrisk/structs"
	"fmt"
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

func GetGeocodes (docID string) LatLongOfBusiness {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}

	var document LatLongOfBusiness

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetGeocodes ERROR:")
		log.Error(err)
	}
	return document
}

func GetSortedArrayFromMongo (docID string,agentID string) []string {
	var resStringArr []string
	collectionName := shashankMongo.DatabaseName.Collection("cluster")
	cursor, err := collectionName.Find(shashankMongo.CtxForDB, bson.M{"bybid":docID})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(shashankMongo.CtxForDB, &episodes); err != nil {
		log.Fatal(err)
	}
	newEpisodeMap:=episodes[0]["sortedArrayWithAgent"]

	for _, value := range newEpisodeMap.(primitive.A) {
		for innerKey, innerValue := range value.(primitive.M) {
			if innerKey==agentID {
				for _, innerMostValue := range innerValue.(primitive.A) {
					resStringArr=append(resStringArr,fmt.Sprintf("%v", innerMostValue))
				}
			}
		}

	}
	return resStringArr
}

func GetSortedArrayOfIdsObjMongo(docID string,agentID string) []DeliveryWithTimeAndDistance {
	collectionName := shashankMongo.DatabaseName.Collection("cluster")
	filter := bson.M{"bybid": docID}

	var document ExtractTimeAndDistanceFromMongo
	var result []DeliveryWithTimeAndDistance

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetGeocodes ERROR:")
		log.Error(err)
	}

	for _,v:=range document.DeliveryDetailObj{
		if v.AgentID==agentID {
			result=v.ArrayOfDeliveryDetail
		}
	}

	return result
}

func DeleteDeliveryFromMongo(docID string)int64{
	collectionName := shashankMongo.DatabaseName.Collection("cluster")
	filter := bson.M{"bybid": docID}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$unset":bson.M{"currentClusterArr": 1,"sortedArrayWithAgent":1}})
	if err != nil {
		log.Error("DeleteDeliveryFromMongo ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func SaveSheetIdToMongo(docID string, sheetID string, link string) int64{
	
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$set":bson.M{"sheetID": sheetID, "sheetLink":link}})
	if err != nil {
		log.Error("SaveSheetIdToMongo ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}