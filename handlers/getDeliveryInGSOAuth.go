
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/bybrisk/delivery-api/data"
	"golang.org/x/oauth2"
	"encoding/base64"
	//"fmt"
	"golang.org/x/oauth2/google"
)

// swagger:route GET /delivery/print/{businessID} delivery printOrdersToSheet
// Google oauth consent screen to print data to google sheet.
//
// responses:
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) PrintOrdersToSheetOAuth (w http.ResponseWriter, r *http.Request) {
	
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://developers.bybrisk.com/delivery/callback",
		ClientID:     "113188653176-fjoovrjckjns6hk9p9nunnp3677omhb3.apps.googleusercontent.com",
		ClientSecret: "C2b3yeljmmSW-rn5WEGJ17kl",
		Scopes:       []string{"https://www.googleapis.com/auth/spreadsheets"},
		Endpoint:     google.Endpoint,
	}

	p.l.Println("Handle GET request -> delivery-api Module")

	vars := mux.Vars(r)
	id := vars["businessID"]

	Data := `{
		"id": "`+id+`"
	}`

	sEnc := base64.StdEncoding.EncodeToString([]byte(Data))

	oauthStateString := sEnc

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

var (
	googleOauthConfig *oauth2.Config
)