package main

import (
	"log"
	"net/http"

	"github.com/bear1278/pocket-bot/config"
	"github.com/bear1278/pocket-bot/src/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
	webhook := tgbotapi.NewWebhook(cfg.WebhookUrl)
	_, err = bot.SetWebhook(webhook)
	if err != nil {
		log.Fatalf("Error init webhook: %s", err)
	}
	handler := handlers.NewHandler(bot)
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.ListenForWebhook("/telegram/webhook")

	go func() {
		if err = http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			handler.HandleCommand(update)
		} else {
			handler.HandleMessage(update)
		}
	}
}
