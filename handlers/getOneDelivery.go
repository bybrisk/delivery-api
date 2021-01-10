
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route GET /delivery/one/{id} delivery getOneDelivery
// Get summerized details of a delivery by its ID.
//
// responses:
//	200: getSingleDeliveryDetail
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) GetSingleDelivery(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> dlivery-api Module")
	
	vars := mux.Vars(r)
	id := vars["id"]

	lp := data.GetOneDelivery(id)

	err := lp.GetAllAgentsResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}