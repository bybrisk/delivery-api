package data

func GetData (docID string) *BusinessAccountResponse {
	account := getBusinessAccount(docID)
	return account
}

func AddData (d *BusinessAccountRequest) *BusinessAccountPostSuccess{
	//save data to database and return ID
	id := createBusinessAccount(d)

	//get and set ProfileConfig based on Business Plan
	res:=getProfileConfig(d)
	_ = setProfileConfig(res,id)

	//set deliveryPending and deliveryDelivered
	_ = setDeliveryStats(id)

	//sending response
	var response = BusinessAccountPostSuccess{
		BybID: id,
		Message: "200_OK_SUCCESS",
	}

	return &response
}

func UpdateData (d *UpdateBusinessAccountRequest) *BusinessAccountPostSuccess {
	res := updateBusinessAccount(d)

	var response BusinessAccountPostSuccess
	//sending response
	if res == 1 {
		response = BusinessAccountPostSuccess{
			BybID: d.BybID,
			Message: "Update Done Successfully",
		}
	}

	return &response

}