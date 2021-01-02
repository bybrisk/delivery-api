
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route GET /delivery/{id} delivery getDelivery
// Get details of a existing delivery by Delivery ID
//
// responses:
//	200: deliveryGetResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) GetDelivery(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> Delivery-api Module")
	
	vars := mux.Vars(r)
	id := vars["id"]
	dateOfDelivery := vars["dateOfDelivery"]

	lp := data.GetDelivery(id,dateOfDelivery)

	err := lp.ToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}