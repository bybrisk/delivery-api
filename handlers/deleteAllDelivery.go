package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route GET /delivery/delete/all/{businessID} delivery deleteAllDelivery
// Delete all deliveries by its BusinessID (BybID).
//
// responses:
//	200: deleteAllDeliveryDetail
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) DeleteAllDelivery(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> delivery-api Module")
	
	vars := mux.Vars(r)
	id := vars["businessID"]

	lp := data.DeleteAllDeliveryByBybID(id)

	err := lp.DeleteAllDeliveryPostSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}