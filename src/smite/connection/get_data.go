package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	G "smite/general"
)

type apiError struct {
	What string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("at %s", e.What)
}

// DataHandler - handles getting fresh data from the API
type DataHandler interface {
	UpdateSession() (*Credentials, error)
	GetSession() (string, error)
	TestSession() (bool, error)
	GetDataUsed() (*[]byte, error)
	GetDemoDetails() (*[]byte, error)
	GetEsportsProLeagueDetails() (*[]byte, error)
	GetFriends() (*[]byte, error)
	GetGodRanks() (*[]byte, error)
	GetGods() (*[]byte, error)
	GetGodRecommendedItems() (*[]byte, error)
	GetItems() (*[]byte, error)
	GetMatchDetails() (*[]byte, error)
	GetMatchPlayerDetails() (*[]byte, error)
	GetMatchIdsByQueue() (*[]byte, error)
	GetLeagueLeaderboard() (*[]byte, error)
	GetLeagueSeasons() (*[]byte, error)
	GetMatchHistory() (*[]byte, error)
	GetPlayer() (*[]byte, error)
	GetPlayerStatus() (*[]byte, error)
	GetQueueStats() (*[]byte, error)
	GetTeamDetails() (*[]byte, error)
	GetTeamMatchHistory() (*[]byte, error)
	GetTeamPlayers() (*[]byte, error)
	GetTopMatches() (*[]byte, error)
	SearchTeams() (*[]byte, error)
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

// Ping - A quick way of validating access to the Hi-Rez API.
func Ping() (*[]byte, error) {
	// /ping[ResponseFormat]
	return buildURL("ping", nil, bMethod()).callAPIAndGetBody(), nil
}

// UpdateSession - Updates the session if it is expires and returns caller
func (credentials *Credentials) UpdateSession() (*Credentials, error) {
	// /createsession[ResponseFormat]/{developerId}/{signature}/{timestamp}
	// Check if current session is valid
	if valid, e := credentials.TestSession(); e == nil && !valid {
		body := buildURL("createsession", credentials, bMethod(), bCredentials(), bTimestamp()).callAPIAndGetBody()
		// Put into struct
		var sj sessionjson
		json.Unmarshal(*body, &sj)
		// Update credentials and save to file
		credentials.SessionID = sj.SessionID
		credentials.saveCredentials()
	}
	return credentials, nil
}

// GetSession - Gets a new session and puts it in credentials
func (credentials *Credentials) GetSession() (string, error) {
	ret, e := credentials.UpdateSession()
	return ret.SessionID, e
}

// TestSession - Tests that the session is valid, returns true if session is valid, false if expired
func (credentials *Credentials) TestSession() (bool, error) {
	// /testsession[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}
	body := buildURL("testsession", credentials, bMethod(), bCredentials(), bSession(), bTimestamp()).callAPIAndGetBody()
	// If response body contains "Invalid" then session has expired
	return !bytes.Contains(*body, []byte("Invalid")), nil
}

// GetDataUsed - returns API Developer daily usage limits and the current status against those limits.
func (credentials *Credentials) GetDataUsed() (*[]byte, error) {
	// /getdataused[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}
	body := buildURL("getdataused", credentials, bMethod(), bCredentials(), bSession(), bTimestamp()).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetDemoDetails - returns information regarding a particular match. Rarely used in lieu of getmatchdetails().
func (credentials *Credentials) GetDemoDetails(matchID int) (*[]byte, error) {
	// /getdemodetails[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{match_id}
	body := buildURL("getdemodetails", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(matchID)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetEsportsProLeagueDetails - Returns the matchup information for each matchup for the current eSports Pro League season. An important return value is
// “match_status” which represents a match being scheduled (1), in-progress (2), or complete (3)
func (credentials *Credentials) GetEsportsProLeagueDetails() (*[]byte, error) {
	// /getesportsproleaguedetails[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}
	body := buildURL("getesportsproleaguedetails", credentials, bMethod(), bCredentials(), bSession(), bTimestamp()).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetFriends - returns the Smite User names of each of the player’s friends.
func (credentials *Credentials) GetFriends(player interface{}) (*[]byte, error) {
	// /getfriends[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{player}
	var body *[]byte
	switch player.(type) {
	case string, int:
		body = buildURL("getfriends", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(player)).callAPIAndGetBody()
	default:
		return nil, &apiError{"player input is of wrong type. Should be of string or int."}
	}
	fmt.Println(string(*body))
	return body, nil
}

// GetGodRanks - returns the Rank and Worshippers value for each God a player has played.
func (credentials *Credentials) GetGodRanks(player interface{}) (*[]byte, error) {
	// /getgodranks[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{player}
	var body *[]byte
	switch player.(type) {
	case string, int:
		body = buildURL("getgodranks", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(player)).callAPIAndGetBody()
	default:
		return nil, &apiError{"player input is of wrong type. Should be of string or int."}
	}
	fmt.Println(string(*body))
	return body, nil
}

// GetGods - gets all current god data and puts it in json file
func (credentials *Credentials) GetGods() (*[]byte, error) {
	// /getgods[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{languageCode}
	body := buildURL("getgods", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bLanguageCode()).callAPIAndGetBody()
	ioutil.WriteFile(G.GodsPath, *body, 777)
	return body, nil
}

// GetGodRecommendedItems - returns the Recommended Items for a particular God.
func (credentials *Credentials) GetGodRecommendedItems(godID int) (*[]byte, error) {
	// /getgodrecommendeditems[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{godid}/{languageCode}
	body := buildURL("getgodrecommendeditems", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(godID), bLanguageCode()).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetItems - gets all current item data and puts it in json file
func (credentials *Credentials) GetItems() (*[]byte, error) {
	// /getitems[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{languagecode}
	body := buildURL("getitems", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bLanguageCode()).callAPIAndGetBody()
	ioutil.WriteFile(G.ItemsPath, *body, 777)
	return body, nil
}

// GetMatchDetails - returns the statistics for a particular completed match.
func (credentials *Credentials) GetMatchDetails() (*[]byte, error) {
	// /getmatchdetails[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{match_id}
	body := buildURL("getmatchdetails", credentials, bMethod(), bCredentials(), bSession(), bTimestamp()).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetMatchPlayerDetails - returns player information for a live match.
func (credentials *Credentials) GetMatchPlayerDetails(matchID int) (*[]byte, error) {
	// /getmatchplayerdetails[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{match_id}
	body := buildURL("getmatchplayerdetails", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(matchID)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetMatchIdsByQueue - lists all Match IDs for a particular Match Queue; useful for API developers interested in constructing data by Queue. To
// limit the data returned, an {hour} parameter was added (valid values: 0 - 23). An {hour} parameter of -1 represents the
// entire day, but be warned that this may be more data than we can return for certain queues. Also, a returned “active_flag”
// means that there is no match information/stats for the corresponding match. Usually due to a match being in-progress,
// though there could be other reasons.
func (credentials *Credentials) GetMatchIdsByQueue(q queue, date, hour string) (*[]byte, error) {
	// /getmatchidsbyqueue[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{queue}/{date}/{hour}
	body := buildURL("getmatchidsbyqueue", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(q), bAny(date), bAny(hour)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetLeagueLeaderboard - returns the top players for a particular league (as indicated by the queue/tier/season parameters).
func (credentials *Credentials) GetLeagueLeaderboard(q queue, t tier, season string) (*[]byte, error) {
	// /getleagueleaderboard[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{queue}/{tier}/{season}
	body := buildURL("getleagueleaderboard", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(q), bAny(t), bAny(season)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetLeagueSeasons - provides a list of seasons (including the single active season) for a match queue.
func (credentials *Credentials) GetLeagueSeasons(q queue) (*[]byte, error) {
	// /getleagueseasons[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{queue}
	body := buildURL("getleagueseasons", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(q)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetMatchHistory - gets recent matches and high level match statistics for a particular player.
func (credentials *Credentials) GetMatchHistory(player interface{}) (*[]byte, error) {
	// /getmatchhistory[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{player}
	var body *[]byte
	switch player.(type) {
	case string, int:
		body = buildURL("getmatchhistory", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(player)).callAPIAndGetBody()
	default:
		return nil, &apiError{"player input is of wrong type. Should be of string or int."}
	}
	fmt.Println(string(*body))
	return body, nil
}

// GetPlayer - returns league and other high level data for a particular player.
func (credentials *Credentials) GetPlayer(playerName string) (*[]byte, error) {
	// /getplayer[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{playerName}
	body := buildURL("getplayer", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(playerName)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetPlayerStatus - returns player status as follows:
func (credentials *Credentials) GetPlayerStatus(player interface{}) (*[]byte, error) {
	// /getplayerstatus[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{player}
	var body *[]byte
	switch player.(type) {
	case string, int:
		body = buildURL("getplayerstatus", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(player)).callAPIAndGetBody()
	default:
		return nil, &apiError{"player input is of wrong type. Should be of string or int."}
	}
	fmt.Println(string(*body))
	return body, nil
}

// GetQueueStats - returns match summary statistics for a (player, queue) combination grouped by gods played.
func (credentials *Credentials) GetQueueStats(player interface{}, q queue) (*[]byte, error) {
	// /getqueuestats[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{player}/{queue}
	var body *[]byte
	switch player.(type) {
	case string, int:
		body = buildURL("getqueuestats", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(player), bAny(q)).callAPIAndGetBody()
	default:
		return nil, &apiError{"player input is of wrong type. Should be of string or int."}
	}
	fmt.Println(string(*body))
	return body, nil
}

// GetTeamDetails - lists the number of players and other high level details for a particular clan.
func (credentials *Credentials) GetTeamDetails(clanID string) (*[]byte, error) {
	// /getteamdetails[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{clanid}
	body := buildURL("getteamdetails", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(clanID)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetTeamMatchHistory - gets recent matches and high level match statistics for a particular clan/team.
func (credentials *Credentials) GetTeamMatchHistory(clanID string) (*[]byte, error) {
	// /getteammatchhistory[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{clanid}
	body := buildURL("getteammatchhistory", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(clanID)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetTeamPlayers - lists the players for a particular clan.
func (credentials *Credentials) GetTeamPlayers(clanID string) (*[]byte, error) {
	// /getteamplayers[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{clanid}
	body := buildURL("getteamplayers", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(clanID)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// GetTopMatches - lists the 50 most watched / most recent recorded matches.
func (credentials *Credentials) GetTopMatches() (*[]byte, error) {
	// /gettopmatches[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}
	body := buildURL("gettopmatches", credentials, bMethod(), bCredentials(), bSession(), bTimestamp()).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}

// SearchTeams - returns high level information for Team names containing the “searchTeam” string.
func (credentials *Credentials) SearchTeams(searchTeam string) (*[]byte, error) {
	// /searchteams[ResponseFormat]/{developerId}/{signature}/{session}/{timestamp}/{searchTeam}
	body := buildURL("searchteams", credentials, bMethod(), bCredentials(), bSession(), bTimestamp(), bAny(searchTeam)).callAPIAndGetBody()
	fmt.Println(string(*body))
	return body, nil
}
