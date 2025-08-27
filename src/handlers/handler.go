package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	storage map[int][]string
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{
		bot:     bot,
		storage: make(map[int][]string),
	}
}

func (h *Handler) HandleMessage(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		return
	}
	log.Printf("Update Message: %s\n", update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello World!")
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := h.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) HandleCommand(update tgbotapi.Update) {
	if !update.Message.IsCommand() {
		return
	}
	id := update.Message.From.ID
	if update.Message.Command() == "setlink" {
		link := update.Message.CommandArguments()
		h.storage[id] = append(h.storage[id], link)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Link was saved")
		_, err := h.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	} else if update.Message.Command() == "links" {
		var sb strings.Builder
		i := int64(1)
		for _, link := range h.storage[id] {
			sb.WriteString(strconv.FormatInt(i, 10))
			sb.WriteString(". ")
			sb.WriteString(link)
			sb.WriteString("\n")
			i++
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, sb.String())
		_, err := h.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
