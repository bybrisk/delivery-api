package data

import (
	"encoding/json"
	"io"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/bybrisk/structs"
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
	ZoneDetailInfo []structs.ZoneInfo `json:"zoneDetailInfo"`
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

/*
11111111111111111111111111111111111111111111111111111111111111111
11111000001110111110111000001110000011100000111000001110111011111
11111011101111011101111011101110111011111011111011111110101111111
11111000001111101011111000001110000011111011111000001110011111111
11111011101111110111111011101110101111111011111111101110101111111
11111000001111110111111000001110111011100000111000001110111011111
11111111111111111111111111111111111111111111111111111111111111111
*/

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