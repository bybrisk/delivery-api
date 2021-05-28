package data

import(
	"net/http"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	//"fmt"
)

func PrintOrderToShareGoogleAPI(docID string){

	//generate login url
	pythonResponse, err := http.Get("https://accounts.google.com/o/oauth2/auth?redirect_uri=http://localhost:8080/docs&response_type=code&client_id=113188653176-fjoovrjckjns6hk9p9nunnp3677omhb3.apps.googleusercontent.com&scope=https://www.googleapis.com/auth/spreadsheets+https://www.googleapis.com/auth/userinfo.email&approval_prompt=force&access_type=offline")
		if err != nil {
			log.Error(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(pythonResponse.Body)
		if err != nil {
			log.Error(err)
		}

		//contentType := pythonResponse.Header.Get("Content-Type")
		//fmt.Println("Content-Type",contentType)
		sb := string(body)
		//fmt.Printf("%s\n",body)
        log.Printf(sb)

		//var betaResponseCrossOrigin PythonClusterAPIResponse
		//err = json.Unmarshal(body, &betaResponseCrossOrigin)
}