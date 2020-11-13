package data

import (
	"encoding/json"
	"io"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/bybrisk/structs"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BusinessAccountRequest struct{
	PicURL string `json: "picurl"`
	UserName string `json: "username" validate:"required"`
	BusinessName string `json: "businessname" validate:"required"`
	Password string `json: "password" validate:"required"` //custom requirement
	City string `json: "city" validate:"required"`
	BusinessPlan string `json: "businessplan" validate:"required"`
}

type BusinessAccountResponse struct{
	PicURL string `json: "picurl"`
	UserName string `json: "username"`
	BusinessName string `json: "businessname"`
	City string `json: "city"`
	BusinessPlan string `json: "businessplan"`
	ProfileConfig structs.ProfileConfig `json:"profileConfiguration"`
	DeliveryPending string `json: "deliveryPending"`
	DeliveryDelivered string `json: "deliveryDelivered"`
	UserID string `json:"BybID"`
	ZoneDetailInfo []structs.ZoneInfo `json:"-"`
}

type BusinessAccountPostSuccess struct {
	BybID string `json:"bybID"`
	Message string `json:"message"`
}

var resultID string

func (d *BusinessAccountRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *BusinessAccountRequest) ToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *BusinessAccountRequest) FromJSON (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *BusinessAccountPostSuccess) ResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

//func GetData () *BusinessAccount {
	//return &deliveryDetail
//}

func AddData (d *BusinessAccountRequest) *BusinessAccountPostSuccess{
	//save data to database and return ID
	id := createBusinessAccount(d)

	//get and set ProfileConfig based on Business Plan
	res:=getProfileConfig(d)
	_ = setProfileConfig(res,id)

	//set deliveryPending and deliveryDelivered
	_ = setDeliveryStats(id)

	//sending response
	var response = BusinessAccountPostSuccess{
		BybID: id,
		Message: "200_OK_SUCCESS",
	}

	return &response
}

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
	update := bson.M{"deliveryPending": "0", "deliveryDelivered":"0"}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("setDeliveryStats ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount
}
