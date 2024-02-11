package utils

import (
	"os"
)

func CreateDir(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

func CheckFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}
