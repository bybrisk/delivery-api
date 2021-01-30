
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/delivery-api/data"
)

// swagger:route GET /delivery/all/history/{agentID} delivery getAllHistoryByAgentID
// Get summerized details of all deliveries (except pending) by its AgentID (BybID).
//
// responses:
//	200: getAllDeliveryDetail
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) GetAllHistoryByAgentID(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> dlivery-api Module")
	
	vars := mux.Vars(r)
	id := vars["agentID"]

	lp := data.GetAgentDeliveryHistory(id)

	err := lp.GetAllDeliveryResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}