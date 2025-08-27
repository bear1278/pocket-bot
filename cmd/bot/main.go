package main

import (
	"github.com/bear1278/pocket-bot/config"
	"github.com/bear1278/pocket-bot/src/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatalf("Error init bot: %s", err)
	}
	handler := handlers.NewHandler(bot)
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		go handler.HandleMessage(update)
		go handler.HandleCommand(update)
	}
}
