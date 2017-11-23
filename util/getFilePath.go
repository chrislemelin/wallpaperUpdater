package util

import (
	"os"
	"strconv"
)

var (
	path          = "pics"
	file          = "temp"
	fileExtension = ".png"
	counter       = 1
)

// GetFilePath gets the next path to be used
func GetFilePath() string {
	thisExists, _ := exists(path)
	if !thisExists {
		os.Mkdir(path, 0755)
	}
	returnValue := path + "/" + file + strconv.Itoa(counter) + fileExtension
	thisExists, _ = exists(returnValue)
	for thisExists {
		counter++
		returnValue = path + "/" + file + strconv.Itoa(counter) + fileExtension
		thisExists, _ = exists(returnValue)
	}
	return returnValue
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
