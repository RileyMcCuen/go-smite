package connection

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	G "smite/general"
)

// Credentials define the credentials of a user accessing the Smite API
type Credentials struct {
	DevID     string `json:"dev_id"`
	AuthKey   string `json:"auth_key"`
	SessionID string `json:"session_id"`
}

// GetCredentialsAndUpdateSession - Gets user's credentials using their CredentialPath
func GetCredentialsAndUpdateSession() *Credentials {
	c := GetCredentials()
	c.GetSession()
	return c
}

// sessionjson - holds SessionID, useful for getting new id from API call
type sessionjson struct {
	SessionID string `json:"session_id"`
}

// GetSignature - creates a signature using provided data and MD5 hash
func (credentials *Credentials) GetSignature(method, timestamp string) (s string) {
	devID, authKey := credentials.DevID, credentials.AuthKey
	s = devID + method + authKey + timestamp
	bytes := md5.Sum([]byte(s))
	s = hex.EncodeToString(bytes[:])
	return s
}

// GetCredentials - Gets user's credentials using their CredentialPath
func GetCredentials() *Credentials {
	data := G.GetDataFromFile(G.CredentialsFile)
	credentials := new(Credentials)
	json.Unmarshal(data, credentials)
	return credentials
}

// saveCredentials - Saves user's credentials using their CredentialPath
func (credentials *Credentials) saveCredentials() {
	if data, e := json.MarshalIndent(credentials, "", " "); G.ErrorHandler(e) {
		G.WriteDataToFile(G.CredentialsFile, data)
	}
}
