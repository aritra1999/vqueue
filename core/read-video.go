package core

import (
	"log"
	"os"
)

func ReadVideo(sourceVideoFilePath string) bool {
	log.Print("Reading file" + sourceVideoFilePath)
	if _, err := os.Stat(sourceVideoFilePath); err == nil {
		log.Print("File exists\n")
		return true
	} else {
		log.Print("File does not exist\n")
		return false
	}
}
