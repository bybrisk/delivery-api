
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route POST /addDelivery delivery addDelivery
// Add a delivery for cluster formation
//
// responses:
//	200: deliveryPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) AddDelivery (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> delivery-api Module")
	delivery := &data.AddDeliveryRequest{}

	err:=delivery.FromJSONToAddDeliveryStruct(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = delivery.ValidateAddDelivery()
	if err!=nil {
		p.l.Println("Validation error in POST request -> delivery-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add delivery to elastic search
	deliveryWithID := data.AddData(delivery)

	//writing to the io.Writer
	err = deliveryWithID.FromAddDeliveryStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}