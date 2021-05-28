
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/bybrisk/delivery-api/data"
	"fmt"
)

// swagger:route GET /delivery/callback delivery printOrdersToSheetCallback
// Print all orders of business to google sheets callback function.
//
// responses:
//	200: printOrderResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) PrintOrdersToSheet (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> delivery-api Module")

	vars := mux.Vars(r)
	id := vars["businessID"]

	fmt.Println(id)

	/*lp := data.GetAgentPendingDelivery(id)

	err := lp.GetAllDeliveryResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}*/
}