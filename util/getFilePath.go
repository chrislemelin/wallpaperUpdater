package util

import (
	"os"
	"time"
)

var (
	path          = "pics"
	fileExtension = ".png"
)

// GetFilePath gets the next path to be used
func GetFilePath() string {
	thisExists, _ := exists(path)
	if !thisExists {
		os.Mkdir(path, 0755)
	}

	t := time.Now()
	return path + "/" + t.Format("2006-01-02-15_04_05") + fileExtension
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
