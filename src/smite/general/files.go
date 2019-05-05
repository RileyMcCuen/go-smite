package general

import (
	"io/ioutil"
)

// File - represents a file than can be easily switched on to get a file path
type File int

// all files
const (
	NoFile          File = -1
	CredentialsFile File = 0
	ItemsFile       File = 1
	GodsFile        File = 2
)

// getFilePathFromFile - gets the file path of the desired file given a File
func getFilePathFromFile(f File) string {
	return [...]string{CredentialsPath, ItemsPath, GodsPath}[f]
}

// GetDataFromFile - gets the data from a file specified by the given File
func GetDataFromFile(f File) []byte {
	filePath := getFilePathFromFile(f)
	file, e := ioutil.ReadFile(filePath)
	ErrorHandler(e)
	return file
}

// WriteDataToFile - writes given data to a file specified by the given File
func WriteDataToFile(f File, data []byte) {
	filePath := getFilePathFromFile(f)
	e := ioutil.WriteFile(filePath, data, 777)
	ErrorHandler(e)
}
