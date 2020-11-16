package data

import (
	"github.com/go-playground/validator/v10"
	"github.com/bybrisk/structs"
)

//post request
type BusinessAccountRequest struct{
	PicURL string `json: "picurl"`
	UserName string `json: "username" validate:"required"`
	BusinessName string `json: "businessname" validate:"required"`
	Password string `json: "password" validate:"required"` //custom requirement
	City string `json: "city" validate:"required"`
	BusinessPlan string `json: "businessplan" validate:"required"`
}

//update request
type UpdateBusinessAccountRequest struct{
	BybID string `json: "bybID" validate:"required"`
	PicURL string `json: "picurl"`
	UserName string `json: "username"`
	BusinessName string `json: "businessname"`
	Password string `json: "password"` //custom requirement
	City string `json: "city"`
}

//get response
type BusinessAccountResponse struct{
	PicURL string `json: "picurl"`
	UserName string `json: "username"`
	BusinessName string `json: "businessname"`
	City string `json: "city"`
	BusinessPlan string `json: "businessplan"`
	ProfileConfig structs.ProfileConfig `json:"profileConfiguration"`
	DeliveryPending string `json: "deliveryPending"`
	DeliveryDelivered string `json: "deliveryDelivered"`
	UserID string `json:"bybID"`
}

//post response
type BusinessAccountPostSuccess struct {
	BybID string `json:"bybID"`
	Message string `json:"message"`
}

func (d *BusinessAccountRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdateBusinessAccountRequest) ValidateUpdateRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}