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
	
	// Status of the payment made by the customer (true or false) - Payment done or not
	//
	PaymentStatus bool `json:"paymentStatus"`
	
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

	//Amount to be colleced in INR ( if payment status is false i.e., Not Done)
	//
	// required: false
	Amount float64 `json:"amount"`

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

	//Delivery Ranking Time (It will be set using an internal algo)
	//
	// required: false
	RankingTime int64 `json:"rankingTime"`

	//TimeStamp of Delivery (It will be set automatically)
	//
	// required: false
	TimeStamp string `json:"timeStamp"`

	// Delivery Observed Distance (No need to set it, it will be done by algo)
	//
	// required: false
	DistanceObserved  float64  `json:"distanceObserved"`
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
	PaymentStatus bool `json:"paymentStatus"`

	// BybID of the business account this delivery is associatd to 
	//
	// required: true
	BybID string `json;"bybID" validate:"required"`

	//Amount to be colleced in INR ( if payment status is false i.e., Not Done)
	//
	// required: false
	Amount float64 `json:"amount"`

	// You donot need to provide the APIKey. It is free rightnow
	//
	APIKey string `json:"apiKey"`

	// You donot need to provide the latitude. It is filled by the API
	//
	Latitude float64 `json:"latitude"`
	
	// You donot need to provide the longitude. It is filled by the API
	//
	Longitude float64 `json:"longitude"`

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

	//Delivery Ranking Time (It will be set using an internal algo)
	//
	// required: false
	RankingTime int64 `json:"rankingTime"`

	//TimeStamp of Delivery (It will be set automatically)
	//
	// required: false
	TimeStamp string `json:"timeStamp"`

	// Delivery Observed Distance (No need to set it, it will be done by algo)
	//
	// required: false
	DistanceObserved  float64  `json:"distanceObserved"`
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

	//TimeStamp of Delivery (It will be set automatically)
	//
	TimeStamp string `json:"timeStamp"`

	//Amount to be colleced in INR ( if payment status is false i.e., Not Done)
	//
	Amount float64 `json:"amount"`
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
type DeliveryResponseBulk struct {
	Hits struct {
		Hits []struct {
			//Date of delivery
			//
			Index  string `json:"_index"`

			//ID of delivery
			//
			ID     string `json:"_id"`

			//Delivery details
			//
			Source struct {
				//Pincode of delivery location
				//
				Pincode         string  `json:"pincode"`

				//API Key used in the delivery
				//
				APIKey          string  `json:"apiKey"`

				//Latitude of delivery location
				//
				Latitude        float64 `json:"latitude"`

				//ClusterID of the cluster this delivery falls into
				//
				ClusterID       string  `json:"clusterID"`

				//AgentID of the agent associated with the delivery
				//
				DeliveryAgentID string  `json:"deliveryAgentID"`

				//Phone number of the customer placing delivery
				//
				Phone           string  `json:"phone"`

				//Name of the customer placing delivery
				//
				CustomerName    string  `json:"CustomerName"`

				//Business ID associated with the delivery
				//
				BybID           string  `json:"BybID"`

				//Weight of Item delivered
				//
				ItemWeight      float64     `json:"itemWeight"`

				//Is payment done or not
				//
				PaymentStatus   bool    `json:"paymentStatus"`

				//Status of Delivery
				//
				DeliveryStatus  string  `json:"deliveryStatus"`

				//Address of delivery
				//
				CustomerAddress string  `json:"CustomerAddress"`

				//Longitude of delivery location
				//
				Longitude       float64 `json:"longitude"`

				//Delivery Ranking Time (It will be set using an internal algo)
				//
				RankingTime int64 `json:"rankingTime"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`

	//Sorted Array of Delivery IDs
	SortedIdArray []DeliveryWithTimeAndDistance `json:"sortedIdArray"`
	SortedIdString []string `json:"SortedIdString"`

}

type DeliveryWithTimeAndDistance struct{
	DeliveryID string `json:"deliveryId"`
	Distance int64 `json:"distance"`
	Time int64 `json:"time"`
}

type MongoStructForTimeAndDistance struct{
	ArrayOfDeliveryDetail []DeliveryWithTimeAndDistance `json:"arrayOfDeliveryDetail"`
	AgentID string `json:"agentID"` 
}

type ExtractTimeAndDistanceFromMongo struct{
	DeliveryDetailObj []MongoStructForTimeAndDistance `json:"deliveryDetailObj"`
}

//post response
type DeliveryPostSuccess struct {
	DeliveryID string `json:"deliveryID"`
	Message string `json:"message"`
}

type UpdateDeliveryDistance struct {
	// DeliveryID of the Delivery in which you want to update the distance 
	//
	// required: true
	DeliveryID string `json;"deliveryID" validate:"required"`

	// Distance travelled by the agent (GPS based) in meters
	//
	// required: true
	Distance float64 `json;"distance" validate:"required"`
}

type DeliveryCountStatus struct {
	DeliveryPending string `json:"deliveryPending"`
	DeliveryDelivered string `json:"deliveryDelivered"`
	DeliveryCancelled string `json: "deliveryCancelled"`
	DeliveryTransit string `json: "deliveryTransit"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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

type LatLongOfBusiness struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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

func (d *UpdateDeliveryDistance) ValidateUpdateDeliveryDistance() error {
	validate := validator.New()
	return validate.Struct(d)
}