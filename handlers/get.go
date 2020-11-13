
package handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/bybrisk/businessAccount-api/data"
)

// swagger:route GET /account businessAccount createBusinessAccount
// Get details of a existing Business Account by ID
//
// responses:
//	200: businessAccountGetResponse
//  422: errorValidation
//  501: errorResponse

func (p *Account) GetAccountDetail(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> businessAccount-api Module")
	
	vars := mux.Vars(r)
	id := vars["id"]

	lp := data.GetData(id)

	err := lp.ToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}