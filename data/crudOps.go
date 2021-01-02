package data

func AddDeliveryWithGeoCode (d *AddDeliveryRequestWithGeoCode) *DeliveryPostSuccess{
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

/*func AddDataWithoutGeocode (d *AddDeliveryWithoutGeocodeRequest) *DeliveryPostSuccess {
	//get geocode using address
	
	//save data to elastic search and return ID

	//update pending deliveries of business account

	//sending response
	
}

func GetDelivery (docID string, dateOfDelivery string) *DeliveryGetResponse{
	delivery := fetchDataFromIndex(docID,dateOfDelivery)
	return delivery
}*/