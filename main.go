package main

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
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

			handleMessage(bot, update.Message)
		}
	},
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch strings.ToLower(message.Text) {
	case "/start":
		handleStart(bot, message)
	case "/goodbye":
		handleGoodbye(bot, message)
	default:
		// Handle other messages here
		reply := tgbotapi.NewMessage(message.Chat.ID, "I don't understand that command.")
		bot.Send(reply)
	}
}

func handleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(message.Chat.ID, "Hello! Welcome to the bot.")
	bot.Send(reply)
}

func handleGoodbye(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(message.Chat.ID, "Goodbye! Have a great day.")
	bot.Send(reply)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
