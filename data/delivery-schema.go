package data

import (
	"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
)

//post request for adding delivery
type AddDeliveryRequest struct{
	CustomerName string `json: "customerName"`
	CustomerAddress string `json: "customerAddress"`
	Phone string `json:"phone"`
	ItemWeight float64 `json:"itemWeight"`
	Pincode string `json:"pincode"`
	PaymentStatus bool `json:"paymentStatus"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	BybID string `json;"bybID"` 
}

//post response
type DeliveryPostSuccess struct {
	DeliveryID string `json:"deliveryID"`
	Message string `json:"message"`
}

func (d *AddDeliveryRequest) ValidateAddDelivery() error {
	validate := validator.New()
	return validate.Struct(d)
}