package data_test

import (
	"testing"
	"fmt"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/delivery-api/data"
)

/*func TestAddDeliveryWithGeoCode(t *testing.T) {
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
}*/

/*func TestAddDeliveryWithoutGeoCode(t *testing.T){
    delivery := &data.AddDeliveryRequestWithoutGeoCode{
		CustomerName : "Mitali Bansal Sharma",
		CustomerAddress : "Shop no.7 new shri ram parisar behind regal homes Awadpuri bhopal, Bhopal-462022, Madhya Pradesh, India",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,	
		BybID : "5ffc9e2d0a550230dc5e26a3",
	}

	res:= data.AddDeliveryWithoutGeoCode(delivery) 
	if res==nil{
		t.Fail()
	}
}*/

/*func TestUpdateDeliveryStatusCO(t *testing.T) {
	update := &data.UpdateDeliveryStatus{
		BybID:"5ffc9e2d0a550230dc5e26a3",
		DeliveryID: "VkMJ-nYBekxrDv_yzbOi",
		DeliveryStatus: "Delivered",
	}

	_= data.UpdateDeliveryStatusCO(update)
}*/

func TestGetAllDeliveries(t *testing.T) {
	res:= data.GetAllDeliveryByBybID("6005bfcb53b27c07c1539ea8")
	fmt.Println(res)
}