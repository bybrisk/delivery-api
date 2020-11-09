package data

import (
	"encoding/json"
	"io"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	//"github.com/bybrisk/structs"
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

//func GetData () *BusinessAccount {
	//return &deliveryDetail
//}

func AddData (d *BusinessAccountRequest) {
	//save data to database and return ID
	id := createBusinessAccount(d)
	fmt.Println(id)
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