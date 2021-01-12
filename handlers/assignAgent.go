
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route POST /delivery/update/agent delivery updateDeliveryAgent
// Assign an Agent to a Delivery.
//
// responses:
//	200: deliveryPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) UpdateDeliveryStatusAPI (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> delivery-api Module")
	delivery := &data.UpdateDeliveryAgent{}

	err:=delivery.FromJSONToUpdateDeliveryAgent(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = delivery.ValidateUpdateDeliveryAgent()
	if err!=nil {
		p.l.Println("Validation error in POST request -> delivery-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add delivery to elastic search
	response := data.UpdateDeliveryAgentCO(delivery)

	//writing to the io.Writer
	err = response.FromAddDeliveryStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}