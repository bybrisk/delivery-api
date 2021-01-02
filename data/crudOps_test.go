package data_test

import (
	"testing"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/delivery-api/data"
)

func TestAddDeliveryWithGeoCode(t *testing.T) {

	delivery := &data.AddDeliveryRequestWithGeoCode{
		CustomerName : "Shashank Sharma",
		CustomerAddress : "A.G Colony, Chetna Samiti, Near Bank of Baroda, Patna, Bihar-800025",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,
		Latitude : 23.4594578,
		Longitude : 77.47784784,	
		BybID : "5ff01d0e2af1c9df782ca7f7",
	}

	res:= data.AddDeliveryWithGeoCode(delivery) 
	if res==nil{
		t.Fail()
	}
}