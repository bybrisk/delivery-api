package data_test

import (
	"testing"
	//"fmt"
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
		BybID : "5ffb2f570a550230dc5e26a0",
	}

	_ = data.AddDeliveryWithGeoCode(delivery)
	
}*/

/*func TestUpdateDeliveryStatus(t *testing.T) {
	update := &data.UpdateDeliveryStatus{
		DeliveryID: "U0Oe8nYBekxrDv_yAbNh",
		DeliveryStatus:"Pending", 
	}

	_=data.UpdateDeliveryStatusCO(update)
}*/

func TestUpdateDeliveryAgent(t *testing.T) {
	update := &data.UpdateDeliveryAgent{
		DeliveryID: "U0Oe8nYBekxrDv_yAbNh",
		DeliveryAgentID:"Amit001", 
	}

	_=data.UpdateDeliveryAgentCO(update)
}

/*func TestInsertDeilveryWithoutGeoCode(t *testing.T) {

	delivery := &data.AddDeliveryRequestWithoutGeoCode{
		CustomerName : "Shashank Prakash",
		CustomerAddress : "A.G Colony, Chetna Samiti, Near Bank of Baroda, Patna, Bihar-800025",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,	
		BybID : "5ffb2f570a550230dc5e26a0",
	}

	_ = data.AddDeliveryWithoutGeoCode(delivery)
	
}*/

/*func TestSearchDocument(t *testing.T) {
	res := data.FetchDeliveryByID("f4c17XYBAsHKFtIUy7yb")
	fmt.Println(res)
}*/