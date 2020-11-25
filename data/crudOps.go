package data

func AddData (d *AddDeliveryRequest) *DeliveryPostSuccess{
	//save data to elastic search and return ID
	insertDataToElastic(d)

	//update pending delivery of business account


	//sending response
	var response = DeliveryPostSuccess{
		DeliveryID: "rfgfdgfdh433443",
		Message: "200_OK_SUCCESS",
	}

	return &response
}