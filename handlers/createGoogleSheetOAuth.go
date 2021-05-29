
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/base64"
)

// swagger:route GET /delivery/print/create/{businessID} delivery createGoogleSheet
// Google oauth consent screen to create a google sheet.
//
// responses:
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) CreateGoogleSheetOAuth (w http.ResponseWriter, r *http.Request) {
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
