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

	//Clusters assigned to the delivery through an internal algo
	//
	// required: false
	ClusterID string `json:"clusterID"`

	//Delivery agent assigned to the delivery using internal algo
	//
	// required: false
	DeliveryAgentID string `json:"deliveryAgentID"`

	//Delivery Status (It will be set to Pending by default)
	//
	// required: false
	DeliveryStatus string `json:"deliveryStatus"`
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

	//Clusters assigned to the delivery through an internal algo
	//
	// required: false
	ClusterID string `json:"clusterID"`

	//Delivery agent assigned to the delivery using internal algo
	//
	// required: false
	DeliveryAgentID string `json:"deliveryAgentID"`

	//Delivery Status (It will be set to Pending by default)
	//
	// required: false
	DeliveryStatus string `json:"deliveryStatus"`
}

//Response of a single Delivery struct
type SingleDeliveryDetail struct {
	// The full Name of the customer
	//
	CustomerName string `json: "customerName"`

	// The full Address of the customer
	//
	CustomerAddress string `json: "customerAddress"`

	// 10 digit mobile number
	//
	Phone string `json:"phone"`
	
	// Weight of the delivery in kg. (By default it is 5kg if not provided. Specify the weight to save on your max weight quota.)
	//
	ItemWeight float64 `json:"itemWeight"`
	
	// Pincode of the customer (specify for better agent allocation and optimization.)
	//
	Pincode string `json:"pincode"`
	
	// Status of the payment made by the customer (true or false)
	//
	PaymentStatus bool `json:"paymentStatus"`

	// latitude of the drop point 
	//
	Latitude float64 `json:"latitude"`
	
	// longitude of the drop point 
	//
	Longitude float64 `json:"longitude"`

	// BybID of the business account this delivery is associatd to 
	//
	BybID string `json;"bybID"`

	//Clusters the delivery is assigned to
	//
	ClusterID string `json:"clusterID"`

	//Delivery agent the delivery is assigned to
	//
	DeliveryAgentID string `json:"deliveryAgentID"`

	//Delivery Status
	//
	DeliveryStatus string `json:"deliveryStatus"`
}

//Update Delivery Status Request
type UpdateDeliveryStatus struct {
	// BybID of the business account this delivery is associatd to 
	//
	BybID string `json;"bybID" validate:"required"`

	// DeliveryID of the Delivery whose status you want to change 
	//
	// required: true
	DeliveryID string `json;"deliveryID" validate:"required"`

	//Delivery Status ( Pending | Transit | Cancelled | Delivered )
	//
	// required: true
	DeliveryStatus string `json:"deliveryStatus" validate:"required"`
}

//Update Delivery Agent Request
type UpdateDeliveryAgent struct {
	// DeliveryID of the Delivery in which you want to assign the agent 
	//
	// required: true
	DeliveryID string `json;"deliveryID" validate:"required"`

	//BybID of the respective agent
	//
	// required: true
	DeliveryAgentID string `json:"deliveryAgentID" validate:"required"`
}

//get all deliveries Response struct
/*type DeliveryResponseBulk struct {
	// Array of deliveries
	//
	Result []SingleDeliveryDetail `json:"result"`

	// BybID of business
	//
	BusinessID string `json:"businessid"`

}*/

type DeliveryResponseBulk struct {
	Hits struct {
		Hits []struct {
			Index  string `json:"_index"`
			ID     string `json:"_id"`
			Source struct {
				Pincode         string  `json:"pincode"`
				APIKey          string  `json:"apiKey"`
				Latitude        float64 `json:"latitude"`
				ClusterID       string  `json:"clusterID"`
				DeliveryAgentID string  `json:"deliveryAgentID"`
				Phone           string  `json:"phone"`
				CustomerName    string  `json:"CustomerName"`
				BybID           string  `json:"BybID"`
				ItemWeight      int     `json:"itemWeight"`
				PaymentStatus   bool    `json:"paymentStatus"`
				DeliveryStatus  string  `json:"deliveryStatus"`
				CustomerAddress string  `json:"CustomerAddress"`
				Longitude       float64 `json:"longitude"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

//post response
type DeliveryPostSuccess struct {
	DeliveryID string `json:"deliveryID"`
	Message string `json:"message"`
}

type DeliveryCountStatus struct {
	DeliveryPending string `json:"deliveryPending"`
	DeliveryDelivered string `json:"deliveryDelivered"`
	DeliveryCancelled string `json: "deliveryCancelled"`
	DeliveryTransit string `json: "deliveryTransit"`
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

func (d *UpdateDeliveryStatus) ValidateUpdateDeliveryStatus() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdateDeliveryAgent) ValidateUpdateDeliveryAgent() error {
	validate := validator.New()
	return validate.Struct(d)
}