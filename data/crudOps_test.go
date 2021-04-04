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
		ItemWeight : 4,
		Pincode : "800025",
		PaymentStatus : false,
		Latitude : 23.235423,
		Longitude : 77.434573,	
		BybID : "6038bd0fc35e3b8e8bd9f81a",
		Amount: 234,
	}

	res:= data.AddDeliveryWithGeoCode(delivery) 
	if res==nil{
		t.Fail()
	}
	fmt.Println(res)
}*/

/*func TestAddDeliveryWithoutGeoCode(t *testing.T){
    delivery := &data.AddDeliveryRequestWithoutGeoCode{
		CustomerName : "Great Donna",
		CustomerAddress : "Hoshangabad, M.P",
		Phone : "9340212623",
		ItemWeight : 6,
		Pincode : "800025",
		PaymentStatus : true,	
		BybID : "6038bd0fc35e3b8e8bd9f81a",
		Amount: 235,
	}

	res:= data.AddDeliveryWithoutGeoCode(delivery) 
	if res==nil{
		t.Fail()
	}
	fmt.Println(res)
}*/

/*func TestUpdateDeliveryStatusCO(t *testing.T) {
	update := &data.UpdateDeliveryStatus{
		BybID:"6017ae0e5b706f00e918d040",
		DeliveryID: "J4aIlHcBMbaQ18HIT5uq",
		DeliveryStatus: "Pending",
	}
	_= data.UpdateDeliveryStatusCO(update)
}*/

/*func TestGetAllDeliveries(t *testing.T) {
	res:= data.GetAllDeliveryByBybID("6013bc1aeef443c14c31f250")
	fmt.Println(res)
}*/

/*func TestGetSingleDelivery(t *testing.T) {
	res:= data.GetOneDelivery("bghPQ3cBtAErZoYVdURZ")
	fmt.Println(res)
}*/

func TestGetPendingDeliveryByAgentID(t *testing.T){
	res:= data.GetAgentPendingDelivery("6038bd71c35e3b8e8bd9f81c")
	fmt.Println(res)
}

/*func TestGetDeliveryHistory(t *testing.T) {
	res:= data.GetAgentDeliveryHistory("601401c24b06c2a9342b3017")
	fmt.Println(res)
}*/

/*func TestUpdateDeliveryDistaneCO(t *testing.T){
	update := &data.UpdateDeliveryDistance{
		DeliveryID: "U41vYncBMpywLSXAVWAa",
		Distance: 340,
	}
	_=data.UpdateDeliveryDistanceCO(update)
}*/