package data_test

import (
	"testing"
	//"fmt"
	"github.com/bybrisk/delivery-api/data"
)

func TestInsertDeilveryWithGeoCode(t *testing.T) {

	delivery := &data.AddDeliveryRequestWithGeoCode{
		CustomerName : "charan Prasad Parihar",
		CustomerAddress : "A.G Colony, Chetna Samiti, Near Bank of Baroda, Patna, Bihar-800025",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,
		Latitude : 23.4594578,
		Longitude : 77.47784784,	
		BybID : "5fff310b00a7284754cd9b57",
	}

	_ = data.AddDeliveryWithGeoCode(delivery)
	
}

/*func TestUpdateDeliveryStatus(t *testing.T) {
	update := &data.UpdateDeliveryStatus{
		BybID:"5ffc9e2d0a550230dc5e26a3",
		DeliveryID: "VEP4-XYBekxrDv_ym7MC",
		DeliveryStatus: "Pending",
	}

	_=data.UpdateDeilveryStatusES(update)
}*/

/*func TestUpdateDeliveryAgent(t *testing.T) {
	updatFetchAllDeliveryES

/*func TestInsertDeilveryWithoutGeoCode(t *testing.T) {
FetchAllDeliveryESincode : "800025",
		PaymentStatus : true,	
		BybID : "5ffc9e2d0a550230dc5e26a3",
	}

	_ = data.AddDeliveryWithoutGeoCode(delivery)
	
}*/

/*func TestSearchDocument(t *testing.T) {
	res := data.FetchDeliveryByID("f4c17XYBAsHKFtIUy7yb")
	fmt.Println(res)
}*/

/*func TestGetAllDeliveries(t *testing.T){
	_= data.FetchAllDeliveryES("5ffc9e2d0a550230dc5e26a3")

}*/