package data

import (
	"encoding/json"
	"io"
	"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	//"github.com/shashank404error/shashankMongo"
)

type BusinessAccountRequest struct{
	PicURL string `json: "picurl"`
	UserName string `json: "username" validate:"required"`
	BusinessName string `json: "businessname" validate:"required"`
	Password string `json: "password" validate:"required"` //custom requirement
	City string `json: "city" validate:"required"`
	BusinessPlan string `json: "businessplan" validate:"required"`
}


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
	
}