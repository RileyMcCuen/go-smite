package apiconnection

import (
	"fmt"
	G "smite/general"
)

type url struct {
	sURL        string
	method      string
	timestamp   string
	credentials *Credentials
}

const (
	// baseSURL - the string url of the smite api
	baseSURL string = "http://api.smitegame.com/smiteapi.svc"
)

var (
	// baseURL - the url that will be used to build all urls
	baseURL = url{sURL: baseSURL, method: "", timestamp: ""}
)

func newURL(method string, credentials *Credentials) *url {
	ret := new(url)
	ret.sURL = baseSURL
	ret.method = method
	ret.timestamp = G.GetTime()
	ret.credentials = credentials
	return ret
}

// bmethod - returns a function that adds the method related info to a url
func bmethod() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s%s", u.sURL, u.method, G.APIResponseFormat.String())
	}
}

// bcredentials - returns a function that adds the credentials related info to a url
func bcredentials() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s/%s", u.sURL, u.credentials.DevID, u.credentials.GetSignature(u.method, u.timestamp))
	}
}

// bsession - returns a function that adds the session related info to a url
func bsession() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s", u.sURL, u.credentials.SessionID)
	}
}

// btimestamp - returns a function that adds the timestamp to a url
func btimestamp() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s", u.sURL, u.timestamp)
	}
}

// bLanguageCode - returns a function that adds the language code to a url
func bLanguageCode() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s", u.sURL, G.APILanguageCode.String())
	}
}

// buildURL - builds a url given the necessary data
func buildURL(method string, credentials *Credentials, builders ...func(u *url)) *url {
	ret := newURL(method, credentials)
	for _, b := range builders {
		b(ret)
	}
	return ret
}

// buildURLStruct - builds the url given the necessary data and returns a struct
func buildURLStruct(method string, credentials *Credentials) *url {
	switch method {
	case "ping":
		return buildURL("ping", nil, bmethod())
	case "createsession":
		return buildURL("createsession", credentials, bmethod(), bcredentials(), btimestamp())
	case "testsession":
		return buildURL("testsession", credentials, bmethod(), bcredentials(), bsession(), btimestamp())
	case "getitems":
		return buildURL("getitems", credentials, bmethod(), bcredentials(), bsession(), btimestamp(), bLanguageCode())
	case "getgods":
		return buildURL("getgods", credentials, bmethod(), bcredentials(), bsession(), btimestamp(), bLanguageCode())
	default:
		return nil
	}
}

// BuildURLString - builds a url given the necessary data
func BuildURLString(method string, credentials *Credentials) string {
	return buildURLStruct(method, credentials).sURL
}
