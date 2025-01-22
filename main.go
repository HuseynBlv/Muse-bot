package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot_token, youtube_api_key := LoadConfig()

	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil {
		log.Fatalf("Error creating new bot: %v", err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Send a search term to find a music, or type /menu to see the available commands.")
				bot.Send(msg)
				continue
			}

			if update.Message.Text == "/menu" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose an option:")
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Download Audio File"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			}

			if update.Message.Text == "Download Audio File" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Send a name to download the audio file.")
				bot.Send(msg)
				continue
			}

			if update.Message.Text != "" && update.Message.Text != "/start" && update.Message.Text != "/menu" && update.Message.Text != "Download Audio File" && update.Message.Text != "Youtube Link" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please wait a little bit👽⏳")
				bot.Send(msg)
				videoURL := YTurl(youtube_api_key, update.Message.Text)
				audioFile := DownloadAudio(videoURL, youtube_api_key, update.Message.Text)

				if audioFile != "" {
					audio := tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(audioFile+".mp3"))
					bot.Send(audio)

					os.Remove(audioFile + ".mp3")
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error downloading audio.")
					bot.Send(msg)
				}
			}

			if update.Message.Text == "Youtube Link" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Send a name to get link for the audio file.")
				bot.Send(msg)
				continue
			}
			if update.Message.Text != "" && update.Message.Text != "/start" && update.Message.Text != "/menu" && update.Message.Text != "Download Audio File" && update.Message.Text != "Youtube Link" {
				query := update.Message.Text

				response := YT(youtube_api_key, query)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			}
		}
	}
}
