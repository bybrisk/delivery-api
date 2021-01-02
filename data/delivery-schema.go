package data

import (
	"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
)

//post request for adding delivery
type AddDeliveryRequestWithGeoCode struct{
	// The full Name of the customer
	//
	// required: true
	// max length: 1000
	CustomerName string `json: "customerName" validate:"required"`

	// The full Address of the customer
	//
	// required: true
	// example: Address1, Address2, City, Pincode
	CustomerAddress string `json: "customerAddress" validate:"required"`

	// 10 digit mobile number
	//
	// required: true
	// max length: 10
	Phone string `json:"phone" validate:"required"`
	
	// Weight of the delivery in kg. (By default it is 5kg if not provided. Specify the weight to save on your max weight quota.)
	//
	// required: true
	ItemWeight float64 `json:"itemWeight" validate:"required"`
	
	// Pincode of the customer (specify for better agent allocation and optimization.)
	//
	Pincode string `json:"pincode" validate:"required"`
	
	// Status of the payment made by the customer (true or false)
	//
	PaymentStatus bool `json:"paymentStatus" validate:"required"`
	
	// Specify the latitude of the drop point (through your application) 
	//
	Latitude float64 `json:"latitude" validate:"required"`
	
	// Specify the longitude of the drop point (through your application) 
	//
	Longitude float64 `json:"longitude" validate:"required"`

	// BybID of the business account this delivery is associatd to 
	//
	BybID string `json;"bybID" validate:"required"` 
}

//post request for adding delivery without geocode
type AddDeliveryWithoutGeocodeRequest struct{
	CustomerName string `json: "customerName" validate:"required"`
	CustomerAddress string `json: "customerAddress" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	ItemWeight float64 `json:"itemWeight" validate:"required"`
	Pincode string `json:"pincode" validate:"required"`
	PaymentStatus bool `json:"paymentStatus" validate:"required"`
	Latitude float64 `json:"-"`
	Longitude float64 `json:"-"`
	BybID string `json;"bybID" validate:"required"` 
}

//post response
type DeliveryPostSuccess struct {
	DeliveryID string `json:"deliveryID"`
	Message string `json:"message"`
}

type DeliveryCountStatus struct {
	DeliveryPending string `json:"deliveryPending"`
	DeliveryDelivered string `json:"deliveryDelivered"`
}

func (d *AddDeliveryRequestWithGeoCode) ValidateAddDelivery() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *AddDeliveryWithoutGeocodeRequest) ValidateAddDeliveryWG() error {
	validate := validator.New()
	return validate.Struct(d)
}