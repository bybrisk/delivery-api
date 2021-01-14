
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route GET /delivery/all/{businessID} delivery getAllDelivery
// Get summerized details of all deliveries by its BusinessID (BybID).
//
// responses:
//	200: getAllDeliveryDetail
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) GetAllDelivery(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> dlivery-api Module")
	
	vars := mux.Vars(r)
	id := vars["businessID"]

	lp := data.GetAllDeliveryByBybID(id)

	err := lp.GetOneDeliveryResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}