package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/telegram-bot-api.v4"
)

var rootCmd = &cobra.Command{
	Use:   "pbot",
	Short: "Telegram Bot with Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("TELE_TOKEN")
		if token == "" {
			log.Fatal("TELE_TOKEN not set in environment variables")
		}

		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			// Handle the incoming message
			handleMessage(bot, update.Message)
		}
	},
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// Add your message handling logic here
	// You can use message.Text to get the text of the message

	// Example: Reply to the user
	reply := tgbotapi.NewMessage(message.Chat.ID, "Hello! I received your message.")
	bot.Send(reply)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
