package connection

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	G "smite/general"
)

// DataHandler - handles getting fresh data from the API
type DataHandler interface {
	UpdateSession() *Credentials
	TestSession() bool
	GetItems()
	GetGods()
}

type apiCaller interface {
	callAPIAndGetBody(u string) *[]byte
}

func (u *url) callAPIAndGetBody() *[]byte {
	response, e := http.Get(u.sURL)
	G.ErrorHandler(e)
	defer response.Body.Close()
	body, e := ioutil.ReadAll(response.Body)
	G.ErrorHandler(e)
	return &body
}

// Ping - Pings the api
func Ping() *[]byte {
	return buildURLStruct("ping", nil).callAPIAndGetBody()
}

// UpdateSession - Updates the session if it is expires and returns caller
func (credentials *Credentials) UpdateSession() *Credentials {
	// Check if current session is valid
	if !credentials.TestSession() {
		body := buildURLStruct("createsession", credentials).callAPIAndGetBody()
		// Put into struct
		var sj sessionjson
		json.Unmarshal(*body, &sj)
		// Update credentials and save to file
		credentials.SessionID = sj.SessionID
		credentials.saveCredentials()
	}
	return credentials
}

// GetSession - Gets a new session and puts it in credentials
func (credentials *Credentials) GetSession() string {
	return credentials.UpdateSession().SessionID
}

// TestSession - Tests that the session is valid, returns true if session is valid, false if expired
func (credentials *Credentials) TestSession() bool {
	body := buildURLStruct("testsession", credentials).callAPIAndGetBody()
	// If response body contains "Invalid" then session has expired
	return !bytes.Contains(*body, []byte("Invalid"))
}

// GetItems - gets all current item data and puts it in json file
func (credentials *Credentials) GetItems() {
	body := buildURLStruct("getitems", credentials).callAPIAndGetBody()
	ioutil.WriteFile(G.ItemsPath, *body, 777)
}

// GetGods - gets all current god data and puts it in json file
func (credentials *Credentials) GetGods() {
	body := buildURLStruct("getgods", credentials).callAPIAndGetBody()
	ioutil.WriteFile(G.GodsPath, *body, 777)
}
