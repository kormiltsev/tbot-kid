package app

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func EditMessage(bot *tgbotapi.BotAPI, id64 int64, messid int, text string) {

	msg := tgbotapi.NewEditMessageText(id64, messid, text)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
