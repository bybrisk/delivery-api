
package handlers

import (
	"net/http"
	//"github.com/gorilla/mux"
	//"github.com/bybrisk/delivery-api/data"
	"golang.org/x/oauth2"
	//"fmt"
	"golang.org/x/oauth2/google"
)

// swagger:route GET /delivery/print/{businessID} delivery printOrdersToSheet
// Google oauth consent screen to print data to google sheet.
//
// responses:
//	200: printOrderResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) PrintOrdersToSheetOAuth (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> delivery-api Module")

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/delivery/print/callback",
		ClientID:     "113188653176-fjoovrjckjns6hk9p9nunnp3677omhb3.apps.googleusercontent.com",
		ClientSecret: "C2b3yeljmmSW-rn5WEGJ17kl",
		Scopes:       []string{"https://www.googleapis.com/auth/spreadsheets"},
		Endpoint:     google.Endpoint,
	}
}