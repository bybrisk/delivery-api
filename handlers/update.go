
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/businessAccount-api/data"
)

// swagger:route POST /account businessAccount createBusinessAccount
// Create a new Business Account
//
// responses:
//	200: businessAccountPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Account) UpdateAccount (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> businessAccount-api Module")
	account := &data.UpdateBusinessAccountRequest{}

	err:=account.FromJSONUpdateRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data 
	err = account.Validate()
	if err!=nil {
		p.l.Println("Validation error in POST request -> businessAccount-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//update account
	accountWithID := data.UpdateData(account)

	//writing to the io.Writer
	err = accountWithID.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}