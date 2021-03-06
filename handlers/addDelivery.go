
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route POST /delivery/create/al delivery addDeliveryWithGeocords
// Add a delivery with Geocode for cluster formation.
//
// responses:
//	200: deliveryPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) AddDeliveryGeoCode (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> delivery-api Module")
	delivery := &data.AddDeliveryRequestWithGeoCode{}

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
	deliveryWithID := data.AddDeliveryWithGeoCode(delivery)

	//writing to the io.Writer
	err = deliveryWithID.FromAddDeliveryStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}