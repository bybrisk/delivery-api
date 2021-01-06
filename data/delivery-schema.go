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
	// required: true
	Pincode string `json:"pincode" validate:"required"`
	
	// Status of the payment made by the customer (true or false)
	//
	// required: true
	PaymentStatus bool `json:"paymentStatus" validate:"required"`
	
	// Specify the latitude of the drop point (through your application) 
	//
	// required: true
	Latitude float64 `json:"latitude" validate:"required"`
	
	// Specify the longitude of the drop point (through your application) 
	//
	// required: true
	Longitude float64 `json:"longitude" validate:"required"`

	// BybID of the business account this delivery is associatd to 
	//
	// required: true
	BybID string `json;"bybID" validate:"required"` 
}

//post request for adding delivery without geocode
type AddDeliveryRequestWithoutGeoCode struct{
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
	// required: true
	Pincode string `json:"pincode" validate:"required"`
	
	// Status of the payment made by the customer (true or false)
	//
	// required: true
	PaymentStatus bool `json:"paymentStatus" validate:"required"`
	
	// You donot need to provide the latitude. It is filled by the API
	//
	Latitude float64 `json:"latitude"`
	
	// You donot need to provide the longitude. It is filled by the API
	//
	Longitude float64 `json:"longitude"`

	// BybID of the business account this delivery is associatd to 
	//
	// required: true
	BybID string `json;"bybID" validate:"required"`

	// You donot need to provide the APIKey. It is free rightnow
	//
	APIKey string `json:"apiKey"`
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

//data structure to access Geocode from Google Map API
type ResponseFromMapAPI struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func (d *AddDeliveryRequestWithGeoCode) ValidateAddDelivery() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *AddDeliveryRequestWithoutGeoCode) ValidateAddDeliveryWG() error {
	validate := validator.New()
	return validate.Struct(d)
}