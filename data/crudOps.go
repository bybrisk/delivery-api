package data

import ( //"log"
		"fmt"
		"time"
		"io/ioutil"
		"net/http"
		"net/url"
		"encoding/json"
		log1 "github.com/sirupsen/logrus"
		"math"
	)

const earthRadius = float64(6371)

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
	var res DeliveryPostSuccess
	var status string 

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
		log1.Error("response from google ERROR : ")
		log1.Error(err)
	}
	var responseObject ResponseFromMapAPI
	json.Unmarshal(responseData, &responseObject)
	if responseObject.Status=="OK"{
		d.Latitude = responseObject.Results[0].Geometry.Location.Lat
		d.Longitude = responseObject.Results[0].Geometry.Location.Lng

		//Restricting address outside 15km radius (haversine)

	} else {
		fmt.Println("GeoCoding Failed")
		//Save The response in db

		//set status to GEOCODING FAILED
		status = "GEOCODING FAILED"

		//return json value.
		res = DeliveryPostSuccess{
			DeliveryID: "",
			Message: status,
		}
		return &res
	}

	//Fetch Pending Delivery
	count:=GetDeliveryFrequency(d.BybID)
	//get latlong of business
	distanceOfDelivery := distanceHaversine(count.Latitude, count.Longitude, d.Latitude, d.Longitude)
	fmt.Println(distanceOfDelivery)
	if (distanceOfDelivery > 15) {
		fmt.Println("GeoCoding Failed! Address out of bound")
		//Save The response in db

		//set status to GEOCODING FAILED
		status = "GEOCODING FAILED"

		//return json value.
		res = DeliveryPostSuccess{
			DeliveryID: "",
			Message: status,
		}
		return &res	
	}

	//Check for distance proximity

	d.DeliveryStatus = "Pending"
	
	t2e2 := time.Now()
	d.RankingTime = t2e2.UnixNano()
	d.TimeStamp = t2e2.Format("2006-Jan-02 3:4:5 PM")

	status = responseObject.Status
	d.APIKey = "API"
	d.DistanceObserved = 0

	//save data to elastic search and return ID
	Id := InsertDeilveryWithoutGeoCode(d)

	//update pending delivery of business account
	_=UpdatePendingDelivery(d.BybID,count.DeliveryPending)


	//sending response
	res = DeliveryPostSuccess{
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

	if (d.DeliveryStatus=="Pending"){
		_=DecreaseTransitDelivery(d.BybID,count.DeliveryTransit)
		_=UpdatePendingDelivery(d.BybID,count.DeliveryPending)
	}
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
	result := FetchAllDeliveryES("BybID",docID)
	
	return &result
}

func GetAgentPendingDelivery(docID string) *DeliveryResponseBulk{

	//Fetch all pending deliveries with agentID
	res := FetchPendingDeliveryByAgentIdES("deliveryAgentID",docID)

	//arrange delivery IDs from the array
	//sortedIDs := GetSortedArray(res)
	sortedIdObjArr := GetSortedArrayOfIdsObjMongo(res.Hits.Hits[0].Source.BybID,res.Hits.Hits[0].Source.ClusterID)
	//fmt.Println(sortedIdObjArr)
	res.SortedIdArray = sortedIdObjArr

	// match the index of the array based on the deliveryIDs

	return res
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

func GetSortedArray(res *DeliveryResponseBulk) []string {
	var minVal float64
	//var sortedRes DeliveryResponseBulk
	var arrOfIds []string
	var arrOfDistance []float64
	var k int
	origin := GetGeocodes(res.Hits.Hits[0].Source.BybID)
	latOrigin:=origin.Latitude
	longOrigin:=origin.Longitude

	for _,_= range res.Hits.Hits{	
		minVal=9999999999999
		for i,hit := range res.Hits.Hits{
			if Find(arrOfIds,hit.ID) == true{
				continue
			} else{
			latFrom:=latOrigin
			lonFrom:=longOrigin
			latTo:=hit.Source.Latitude
			lonTo:=hit.Source.Longitude
			var deltaLat = (latTo - latFrom) * (math.Pi / 180)
			var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)
			
			var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) + 
				math.Cos(latFrom * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
				math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
			var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
			
			distance := earthRadius * c
			if distance < minVal {
				minVal = distance
				k=i
			}
		 }
		}
		arrOfIds=append(arrOfIds,res.Hits.Hits[k].ID)
		arrOfDistance=append(arrOfDistance,minVal)
		latOrigin =res.Hits.Hits[k].Source.Latitude
		longOrigin=res.Hits.Hits[k].Source.Longitude
	}
	return arrOfIds	
}

func Find(slice []string, val string)  bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

func distanceHaversine(latFrom float64,lonFrom float64, latTo float64, lonTo float64) float64{

	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)
			
	var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) + 
				math.Cos(latFrom * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
				math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
	var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
			
	distance := earthRadius * c
	return distance
}

