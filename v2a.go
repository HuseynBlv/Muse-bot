package main

import (
	"fmt"
	"os/exec"
)

func DownloadAudio(videoURL, apiKey, qr string) string {

	outputFile := fmt.Sprintf(title(apiKey, qr))
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "-o", outputFile, videoURL)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error downloading audio:", err)
		return ""
	}
	return outputFile
}
