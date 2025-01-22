package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	bot_token := os.Getenv("BOT_TOKEN")
	youtube_api_key := os.Getenv("YOUTUBE_API_KEY")

	return bot_token, youtube_api_key
}
