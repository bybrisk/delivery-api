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
		CustomerName : "Charan Parihar",
		CustomerAddress : "A.G Colony, Chetna Samiti, Near Bank of Baroda, Patna, Bihar-800025",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : false,
		Latitude : 23.4594578,
		Longitude : 77.47784784,	
		BybID : "60093975e6c847545ac2fdf1",
		Amount: 234,
	}

	res:= data.AddDeliveryWithGeoCode(delivery) 
	if res==nil{
		t.Fail()
	}
}*/

/*func TestAddDeliveryWithoutGeoCode(t *testing.T){
    delivery := &data.AddDeliveryRequestWithoutGeoCode{
		CustomerName : "Latest By all means",
		CustomerAddress : "Shop no.7 new shri ram parisar behind regal homes Awadpuri bhopal, Bhopal-462022, Madhya Pradesh, India",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,	
		BybID : "60093975e6c847545ac2fdf1",
		Amount: 235,
	}

	res:= data.AddDeliveryWithoutGeoCode(delivery) 
	if res==nil{
		t.Fail()
	}
}*/

/*func TestUpdateDeliveryStatusCO(t *testing.T) {
	update := &data.UpdateDeliveryStatus{
		BybID:"600d95c5d72ee5dd5896dd75",
		DeliveryID: "bghPQ3cBtAErZoYVdURZ",
		DeliveryStatus: "Pending-Cancelled",
	}

	_= data.UpdateDeliveryStatusCO(update)
}*/

func TestGetAllDeliveries(t *testing.T) {
	res:= data.GetAllDeliveryByBybID("60093975e6c847545ac2fdf1")
	fmt.Println(res)
}

/*func TestGetSingleDelivery(t *testing.T) {
	res:= data.GetOneDelivery("bghPQ3cBtAErZoYVdURZ")
	fmt.Println(res)
}*/