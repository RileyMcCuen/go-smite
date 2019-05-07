package connection

import (
	"fmt"
	"regexp"
	G "smite/general"
	"strings"
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

// the status of a player
type playerstatus int

const (
	//  Offline
	offline playerstatus = 0
	// In Lobby (basically anywhere except god selection or in game)
	inlobby playerstatus = 1
	// God Selection (player has accepted match and is selecting god before start of game)
	godselection playerstatus = 2
	// In Game (match has started)
	ingame playerstatus = 3
	// Online (player is logged in, but may be blocking broadcast of player state)
	online playerstatus = 4
)

// the id of the game mode
type queue int

// all queues in SMITE
const (
	Conquest5v5         queue = 423
	NoviceQueue         queue = 424
	Conquest            queue = 426
	Practice            queue = 427
	ConquestChallenge   queue = 429
	ConquestRanked      queue = 430
	Domination          queue = 433
	MOTD434             queue = 434 //(use with 465 to get all MOTD matches)
	Arena               queue = 435
	ArenaChallenge      queue = 438
	DominationChallenge queue = 439
	JoustLeague         queue = 440
	JoustChallenge      queue = 441
	Assault             queue = 445
	AssaultChallenge    queue = 446
	Joust3v3            queue = 448
	ConquestLeague      queue = 451
	ArenaLeague         queue = 452
	MOTD465             queue = 465 //(Supports “closing” the Queue by our platform; use with 434)
)

// the tier of the match
type tier int

// all tiers in SMITE
const (
	BronzeV     tier = 1
	BronzeIV    tier = 2
	BronzeIII   tier = 3
	BronzeII    tier = 4
	BronzeI     tier = 5
	SilverV     tier = 6
	SilverIV    tier = 7
	SilverIII   tier = 8
	SilverII    tier = 9
	SilverI     tier = 10
	GoldV       tier = 11
	GoldIV      tier = 12
	GoldIII     tier = 13
	GoldII      tier = 14
	GoldI       tier = 15
	PlatinumV   tier = 16
	PlatinumIV  tier = 17
	PlatinumIII tier = 18
	PlatinumII  tier = 19
	PlatinumI   tier = 20
	DiamondV    tier = 21
	DiamondIV   tier = 22
	DiamondIII  tier = 23
	DiamondII   tier = 24
	DiamondI    tier = 25
	MastersI    tier = 26
)

// bmethod - returns a function that adds the method related info to a url
func bMethod() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s%s", u.sURL, u.method, G.APIResponseFormat.String())
	}
}

// bcredentials - returns a function that adds the credentials related info to a url
func bCredentials() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s/%s", u.sURL, u.credentials.DevID, u.credentials.GetSignature(u.method, u.timestamp))
	}
}

// bsession - returns a function that adds the session related info to a url
func bSession() func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s", u.sURL, u.credentials.SessionID)
	}
}

// btimestamp - returns a function that adds the timestamp to a url
func bTimestamp() func(u *url) {
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

// bAny - returns a function that adds any arguments to the end of the url
func bAny(any interface{}) func(u *url) {
	return func(u *url) {
		u.sURL = fmt.Sprintf("%s/%s", u.sURL, fmt.Sprint(any))
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

// BuildAPIMethod - small function that I helped to use to parse api documentation to make functions quicker
func BuildAPIMethod(name, apicall string) {
	var sbuilder strings.Builder
	sbuilder.WriteString(`buildURL(`)
	method, _ := regexp.Compile("([a-z]+)(?:\\[ResponseFormat\\])")
	methodLowercaseName := method.FindStringSubmatch(apicall)[1]
	sbuilder.WriteString(fmt.Sprintf(`"%s", credentials, bMethod(), `, methodLowercaseName))
	build, _ := regexp.Compile("(?:\\{)([a-zA-Z]*)(?:\\})")
	matches := build.FindAllStringSubmatch(apicall, -1)
	for _, match := range matches {
		m := match[1]
		switch m {
		case "developerId":
			sbuilder.WriteString(fmt.Sprintf(`%s, `, "bCredentials()"))
			break
		case "signature":
			break
		case "session":
			sbuilder.WriteString(fmt.Sprintf(`%s, `, "bSession()"))
			break
		case "timestamp":
			sbuilder.WriteString(fmt.Sprintf(`%s, `, "bTimestamp()"))
			break
		case "languageCode":
			sbuilder.WriteString(fmt.Sprintf(`%s, `, "bLanguageCode()"))
			break
		default:
			sbuilder.WriteString(fmt.Sprintf(`bAny(%s), `, m))
			break
		}
	}
	sl := sbuilder.String()[:sbuilder.Len()-2]
	var finb strings.Builder
	finb.WriteString(sl)
	finb.WriteString(")")
	fmt.Printf(
		`
		func (credentials *Credentials) %s() (*[]byte, error) {
			// %s
			body := %s.callAPIAndGetBody()
			fmt.Println(string(*body))
			return body, nil
		}
		`, name, apicall, finb.String())
}
