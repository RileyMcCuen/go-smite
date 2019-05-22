package general

import (
	"strconv"
	"strings"
	"time"
)

// LanguageCode - code corresponding to the lanuage that the API will respond with
type LanguageCode int

// all possible language codes allowed by the API
const (
	English LanguageCode = 1
	German  LanguageCode = 2
	French  LanguageCode = 3
	Spanish LanguageCode = 7
	// Spanish - Latin American
	SpanishLA  LanguageCode = 9
	Portuguese LanguageCode = 10
	Russian    LanguageCode = 11
	Polish     LanguageCode = 12
	Turkish    LanguageCode = 13
)

func (a LanguageCode) String() string {
	return strconv.Itoa(int(a))
}

// ResponseFormat - int corresponding to the kind of response we will expect from the API
type ResponseFormat int

// All accepted response formats from the API
const (
	json ResponseFormat = 0
	xml  ResponseFormat = 1
)

func (r ResponseFormat) String() string {
	return [...]string{"json", "xml"}[r]
}

// ErrorHandler - handle errors by panicking
func ErrorHandler(e error) bool {
	if e != nil {
		panic(e)
	}
	return true
}

// GetTime - gets time in yyyyMMddHHmmss format
func GetTime() string {
	return time.Now().UTC().Format("20060102150405")
}

// CommonStringToBool - converts commonly used strings to booleans, anything unusual defaults to false
func CommonStringToBool(s string) bool {
	s = strings.ToLower(s)
	if s == "y" || s == "yes" || s == "true" {
		return true
	}
	return false
}
