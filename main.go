package main

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/astravexton/telegram-bot-api"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("You must set BOT_TOKEN")
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.From.ID == 777000 && !update.Message.IsAutomaticForward {
				log.Printf("Message sent as channel from @%s, deleting and banning\n",
					update.Message.SenderChat.UserName)
				ban := tgbotapi.BanChatSenderChatConfig{
					ChatID:       update.Message.Chat.ID,
					SenderChatID: update.Message.SenderChat.ID,
					UntilDate:    int(time.Now().Add(1).Unix()),
				}
				bot.Send(ban)
				delete := tgbotapi.DeleteMessageConfig{
					ChatID:    update.Message.Chat.ID,
					MessageID: update.Message.MessageID,
				}
				bot.Send(delete)
			}
		}
	}
}
