package data

import ("log"
		"fmt"
		"time"
		"io/ioutil"
		"net/http"
		"net/url"
		"encoding/json"
	)

func AddDeliveryWithGeoCode (d *AddDeliveryRequestWithGeoCode) *DeliveryPostSuccess{

	d.DeliveryStatus = "Pending"
	d.DistanceObserved = 0
	t2e2 := time.Now()
	d.RankingTime = t2e2.UnixNano()
	d.TimeStamp = t2e2.Format("2006-Jan-02 3:4:5 PM")
	//save data to elastic search and return ID
	res := InsertDeilveryWithGeoCode(d)

	//Fetch Pending Delivery
	count:=GetDeliveryFrequency(d.BybID)
	//update pending delivery of business account
	_=UpdatePendingDelivery(d.BybID,count.DeliveryPending)

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
	
	t2e2 := time.Now()
	d.RankingTime = t2e2.UnixNano()
	d.TimeStamp = t2e2.Format("2006-Jan-02 3:4:5 PM")

	status := responseObject.Status
	d.APIKey = "API"
	d.DistanceObserved = 0

	//save data to elastic search and return ID
	Id := InsertDeilveryWithoutGeoCode(d)

	//Fetch Pending Delivery
	count:=GetDeliveryFrequency(d.BybID)
	//update pending delivery of business account
	_=UpdatePendingDelivery(d.BybID,count.DeliveryPending)

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
	_ = UpdateDeilveryStatusES(d)

	//Fetch frequency of this status 
	count:=GetDeliveryFrequency(d.BybID)

	if (d.DeliveryStatus=="Transit"){
		_=DecreasePendingDelivery(d.BybID,count.DeliveryPending)
		_=UpdateTransitDelivery(d.BybID,count.DeliveryTransit)
	}
	if (d.DeliveryStatus=="Cancelled"){
		_=DecreaseTransitDelivery(d.BybID,count.DeliveryTransit)
		_=UpdateCancelledDelivery(d.BybID,count.DeliveryCancelled)
	}
	if (d.DeliveryStatus=="Delivered"){
		_=DecreaseTransitDelivery(d.BybID,count.DeliveryTransit)
		_=UpdateDeliveredDelivery(d.BybID,count.DeliveryDelivered)
	}
	if (d.DeliveryStatus=="Pending-Cancelled"){
		_=DecreasePendingDelivery(d.BybID,count.DeliveryPending)
		_=UpdateCancelledDelivery(d.BybID,count.DeliveryCancelled)
	}
	//sending response
	response := DeliveryPostSuccess{
		DeliveryID: d.DeliveryID,
		Message: "Delivery Status Updated",
	}

	return &response
}

func UpdateDeliveryAgentCO(d *UpdateDeliveryAgent) *DeliveryPostSuccess {
	//Update Delivery Status in ES Queue
	_ = UpdateDeilveryAgentES(d)

	//sending response
	response := DeliveryPostSuccess{
		DeliveryID: d.DeliveryID,
		Message: "Delivery Agents Assigned",
	}
	
	return &response
}

func GetAllDeliveryByBybID(docID string) *DeliveryResponseBulk {

	//Fetch all deliveries having similar businessID
	res := FetchAllDeliveryES("BybID",docID)
	
	return &res
}

func GetAgentPendingDelivery(docID string) *DeliveryResponseBulk{

	//Fetch all pending deliveries with agentID
	res := FetchPendingDeliveryByAgentIdES("deliveryAgentID",docID)

	return &res
}

func GetAgentDeliveryHistory(docID string) *DeliveryResponseBulk{

	//Fetch all deliveries of an agent which are not pending
	res := FetchDeliveryHistoryByAgentIdES("deliveryAgentID",docID)

	return &res
}

func UpdateDeliveryDistanceCO(d *UpdateDeliveryDistance) *DeliveryPostSuccess {
	//Update Delivery distance in ES Queue
	_ = UpdateDeilveryDistanceES(d)

	//sending response
	response := DeliveryPostSuccess{
		DeliveryID: d.DeliveryID,
		Message: "Delivery Distance Updated",
	}
	
	return &response
}