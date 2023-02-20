package core

import (
	"log"
	"os"
)

func readVideo(filename string) {
	video, err := os.Open("input/video.mp4")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(video.)
}
