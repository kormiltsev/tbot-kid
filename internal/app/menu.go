package app

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Menu(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
	t := "loading..."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, t)
	a, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
	<-time.After(time.Second * 1)
	t = "Actual commands for SuperUsers:\n/kidaddsuperuser_##########\n/kidadduser_##########\n/kidlistofusers\n/kidsleep - start sleep\n/kidwu - stop sleep\n/kidaskcatalog - catalog"
	msgr := tgbotapi.NewEditMessageText(update.Message.Chat.ID, a.MessageID, t)
	if _, err = bot.Send(msgr); err != nil {
		log.Panic(err)
	}

}
