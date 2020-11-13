package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	//"github.com/bybrisk/businessAccount-api/data"
)

type Account struct {
 l *log.Logger
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
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

func (p *Account) UpdateAccount (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request -> businessAccount-api Module")
	vars := mux.Vars(r)
	id := vars["id"]
	p.l.Println("Update businessAccount content with id",id)
}