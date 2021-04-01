package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/delivery-api/data"
)

/*func TestInsertDeilveryWithGeoCode(t *testing.T) {

	delivery := &data.AddDeliveryRequestWithGeoCode{
		CustomerName : "charan Prasad Parihar",
		CustomerAddress : "A.G Colony, Chetna Samiti, Near Bank of Baroda, Patna, Bihar-800025",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,
		Latitude : 23.4594578,
		Longitude : 77.47784784,	
		BybID : "6038bd0fc35e3b8e8bd9f81a",
	}

	_ = data.AddDeliveryWithGeoCode(delivery)
	
}*/

/*func TestUpdateDeliveryStatus(t *testing.T) {
	update := &data.UpdateDeliveryStatus{
		BybID:"6038bd0fc35e3b8e8bd9f81a",
		DeliveryID: "F7cng3gBsGM1IID4S8zs",
		DeliveryStatus: "Transit",
	}

	res:=data.UpdateDeilveryStatusES(update)
	fmt.Println(res)
}*/

/*func TestUpdateDeilveryAgentES(t *testing.T) {
	update := &data.UpdateDeliveryAgent{
		DeliveryID: "F7cng3gBsGM1IID4S8zs",
		DeliveryAgentID: "603b496dc35e3b8e8bd9f83a",
	}

	res:=data.UpdateDeilveryAgentES(update)
	fmt.Println(res)
}*/


/*func TestInsertDeilveryWithoutGeoCode(t *testing.T) {
	FetchAllDeliveryESincode : "800025",
		PaymentStatus : true,	
		BybID : "5ffc9e2d0a550230dc5e26a3",
	}

	_ = data.AddDeliveryWithoutGeoCode(delivery)
	
}*/

/*func TestSearchDocument(t *testing.T) {
	res := data.FetchDeliveryByID("F7cng3gBsGM1IID4S8zs")
	fmt.Println(res)
}*/

/*func TestGetAllDeliveries(t *testing.T){
	res:= data.FetchAllDeliveryES("BybID","6038bd0fc35e3b8e8bd9f81a")
	fmt.Println(res)
}*/

func TestDeliveryDistance(t *testing.T){
	update := &data.UpdateDeliveryDistance{
		DeliveryID: "F7cng3gBsGM1IID4S8zs",
		Distance: 12345,
	}
	res:=data.UpdateDeilveryDistanceES(update)
	fmt.Println(res)
}