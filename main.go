package main

import (
	"vqueue/core"
)

func main() {
	fileName := "demo.mp4"
	sourceVideoFilePath := "input/" + fileName
	outputPath := core.GenerateOutputStructure()
	if core.ReadVideo(sourceVideoFilePath) {
		core.ProcessVideo(sourceVideoFilePath, outputPath)
	}
}
