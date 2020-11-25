package data

func AddData (d *AddDeliveryRequest) *DeliveryPostSuccess{
	//save data to elastic search and return ID
	id:=insertDataToElastic(d)

	//update pending delivery of business account


	//sending response
	var response = DeliveryPostSuccess{
		DeliveryID: id,
		Message: "200_OK_SUCCESS",
	}

	return &response
}