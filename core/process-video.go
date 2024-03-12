package core

import (
	"log"
	"os/exec"
)

func generateTimeLineThumbnails(sourceVideoFile string, outputPath string) {
	log.Print("Generating timeline thumbnails for " + sourceVideoFile)
}

func generateVideoThumbnails(sourceVideoFile string, outputPath string) {
	log.Print("Generating video thumbnails for " + sourceVideoFile)
}

func generateDifferentResolutionVideos(sourceVideoFile string, outputPath string) {
	log.Print("Generating video thumbnails for " + sourceVideoFile)
	resolutions := [
		{
			X: 480,
			Y: 854
		},
		{
			X: 720,
			Y: 1200
		},
		{
			X: 1920,
			Y: 1080
		}
	]

	for resolution := range resolutions {
		// output file paths
		output1 := outputPath + resolution.X + "x" + resolution.Y + ".mp4"
		// ffmpeg command to generate 480p output
		cmd := exec.Command("ffmpeg", "-i", sourceVideoFile, "-vf", "scale=854:480", "-c:a", "copy", output1)
		cmd.Run()
	}


}

func ProcessVideo(sourceVideoFile string, outputPath string) {
	generateDifferentResolutionVideos(sourceVideoFile, outputPath)
	generateVideoThumbnails(sourceVideoFile, outputPath)
	generateTimeLineThumbnails(sourceVideoFile, outputPath)
}
