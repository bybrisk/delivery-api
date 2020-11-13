// Package classification of BusinessAccount API
//
// Documentation for BusinessAccount API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/bybrisk/businessAccount-api/data"
)

type Account struct {
 l *log.Logger
}

func NewAccount(l *log.Logger) *Account{
	return &Account{l}
}

/*func (p *Account) GetAccountDetail(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> businessAccount-api Module")
	lp := data.GetData()

	err := lp.ToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}*/

func (p *Account) AddNewAccount (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> businessAccount-api Module")
	account := &data.BusinessAccountRequest{}

	err:=account.FromJSON(r.Body)
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
	data.AddData(account)
}

func (p *Account) UpdateAccount (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request -> businessAccount-api Module")
	vars := mux.Vars(r)
	id := vars["id"]
	p.l.Println("Update businessAccount content with id",id)
}