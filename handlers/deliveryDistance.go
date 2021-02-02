
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route POST /delivery/update/distance delivery updateDeliveryDistance
// Update distance travelled by an Agent assigned to a Delivery. (GPS Distance)
//
// responses:
//	200: deliveryPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) UpdateDeliveryDistanceAPI (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> delivery-api Module")
	delivery := &data.UpdateDeliveryDistance{}

	err:=delivery.FromJSONToUpdateDeliveryDistance(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = delivery.ValidateUpdateDeliveryDistance()
	if err!=nil {
		p.l.Println("Validation error in POST request -> delivery-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//Update delivery in elastic search
	response := data.UpdateDeliveryDistanceCO(delivery)

	//writing to the io.Writer
	err = response.FromAddDeliveryStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}