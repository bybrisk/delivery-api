package data

import ("log"
		"fmt"
		"io/ioutil"
		"net/http"
		"net/url"
		"encoding/json")

func AddDeliveryWithGeoCode (d *AddDeliveryRequestWithGeoCode) *DeliveryPostSuccess{

	d.DeliveryStatus = "Pending"
	//save data to elastic search and return ID
	res := InsertDeilveryWithGeoCode(d)

	//Fetch Pending Delivery
	count:=GetPendingDelivery(d.BybID)
	//update pending delivery of business account
	_=UpdatePendingDlivery(d.BybID,count.DeliveryPending)

	//sending response
	var response = DeliveryPostSuccess{
		DeliveryID: res,
		Message: "Delivery added to ES Queue",
	}

	return &response
}

func AddDeliveryWithoutGeoCode (d *AddDeliveryRequestWithoutGeoCode) *DeliveryPostSuccess {
	apiKey := "AIzaSyAZDoWPn-emuLvzohH3v-cS_En-u9NSA1A"
	address := url.QueryEscape(d.CustomerAddress)
	url :=  "https://maps.googleapis.com/maps/api/geocode/json?address="+address+"&key="+apiKey
	//get geocode using address
	response, err := http.Get(url)

    if err != nil {
        fmt.Print(err.Error())
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	var responseObject ResponseFromMapAPI
	json.Unmarshal(responseData, &responseObject)
	d.Latitude = responseObject.Results[0].Geometry.Location.Lat
	d.Longitude = responseObject.Results[0].Geometry.Location.Lng
	d.DeliveryStatus = "Pending"
	status := responseObject.Status
	d.APIKey = "API"

	//save data to elastic search and return ID
	Id := InsertDeilveryWithoutGeoCode(d)

	//Fetch Pending Delivery
	count:=GetPendingDelivery(d.BybID)
	//update pending delivery of business account
	_=UpdatePendingDlivery(d.BybID,count.DeliveryPending)

	//sending response
	var res = DeliveryPostSuccess{
		DeliveryID: Id,
		Message: status,
	}
	return &res
}

func GetOneDelivery(docID string) *SingleDeliveryDetail {

	//Fetch the document from elastic search queue
	res := FetchDeliveryByID(docID)

	return &res
}

func UpdateDeliveryStatusCO(d *UpdateDeliveryStatus) *DeliveryPostSuccess {
	//Update Delivery Status in ES Queue
	res := UpdateDeilveryStatusES(d)

	//sending response
	response := DeliveryPostSuccess{
		DeliveryID: res,
		Message: "Delivery Status Updated",
	}

	return &response
}

func UpdateDeliveryAgentCO(d *UpdateDeliveryAgent) *DeliveryPostSuccess {
	//Update Delivery Status in ES Queue
	res := UpdateDeilveryAgentES(d)

	//sending response
	response := DeliveryPostSuccess{
		DeliveryID: res,
		Message: "Delivery Agents Assigned",
	}
	
	return &response
}