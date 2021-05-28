
package data

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"

	"context"
	"reflect"
    "fmt"
    "io/ioutil"
    "net/http"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/delivery/callback",
		ClientID:     "113188653176-fjoovrjckjns6hk9p9nunnp3677omhb3.apps.googleusercontent.com",
		ClientSecret: "C2b3yeljmmSW-rn5WEGJ17kl",
		Scopes:       []string{"https://www.googleapis.com/auth/spreadsheets"},
		Endpoint:     google.Endpoint,
	}
}

func PrintOrderToShareGoogleAPI(docID string, r *http.Request) {
	ctx := context.Background()
	_, token, err := GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		log.Error("PrintOrderToShareGoogleAPI ERROR:")
		log.Error(err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config, token)

	//Google sheet code starts from here
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
            log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // Prints the names and majors of students in a sample spreadsheet:
    // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
    spreadsheetId := "1F1FdtaJJVR2Mc2vfunPV_99Hp6ulq4U5lIKWEo-0qeA"
    readRange := "Sheet1!A2:C30"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve data from sheet: %v", err)
        }

        if len(resp.Values) == 0 {
                fmt.Println("No data found.")
        } else {
                fmt.Println("Name, Major:")
                for _, row := range resp.Values {
                        // Print columns A and E, which correspond to indices 0 and 4.
                        fmt.Printf("%s, %s\n", row[0], row[4])
                }
        }

}

func GetUserInfo(state string, code string) ([]byte, *oauth2.Token, error) {
	if state != oauthStateString {
		return nil, nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, nil ,fmt.Errorf("code exchange failed: %s", err.Error())
	}
	fmt.Println("token = ", reflect.TypeOf(token))

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, token, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config, tok *oauth2.Token) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	return config.Client(context.Background(), tok)
}