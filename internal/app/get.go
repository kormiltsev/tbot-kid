package app

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	st "github.com/kormiltsev/tbot-kid/internal/storage"

	"time"
)

// func AskCatalog(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фильтр:\n/kidaskcatalogall - весь каталог\n")
// 	if _, err := bot.Send(msg); err == nil {
// 		log.Panic(err)
// 	}

// }

// func CatalogDialog(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	t := "За сколько дней выгрузить статистику?"
// 	msgr := tgbotapi.NewMessage(update.Message.Chat.ID, t)
// 	if _, err := bot.Send(msgr); err != nil {
// 		log.Panic(err)
// 	}
// }

func GetRows(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	date2 := time.Now()
	date1 := time.Date(2020, time.December, 31, 21, 0, 0, 0, time.Local)
	msgr := st.GetData("none", date1, date2)
	log.Printf("get rows: ", msgr)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgr)
	if _, err := bot.Send(msg); err == nil {
		log.Panic(err)
	}
}
