package core

import (
	"log"
	"os"
	"time"
)

func generateDirectory(directoryName string) {
	err := os.MkdirAll(directoryName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func GenerateOutputStructure() {
	outputDirectories := [3]string{"video-thumbnail", "timeline-thumbnail", "videos"}
	outputParentDirectory := "output/" + time.Now().Format(time.RFC850)
	log.Print("Creating output files in '" + outputParentDirectory + "'")
	generateDirectory(outputParentDirectory)

	for _, item := range outputDirectories {
		generateDirectory(outputParentDirectory + "/" + item)
	}
}
