
package handlers

import (
	"net/http"
	"github.com/bybrisk/delivery-api/data"
	"encoding/base64"
	"encoding/json"
)

// swagger:route GET /delivery/create/callback delivery createGoogleSheetCallback
// Create google sheets callback function.
//
// responses:
//	200: createGoogleSheetResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) CreateGoogleSheetCallback (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> delivery-api Module")

	state := r.URL.Query().Get("state")

	sDec, _ := base64.StdEncoding.DecodeString(state)

	request := data.GoogleSheetStructDir{}

	json.Unmarshal(sDec, &request)

	lp := data.CreateGoogleSheetCrudOps(request.Id,r)

	err := lp.GoogleSpreadSheetMetaStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}