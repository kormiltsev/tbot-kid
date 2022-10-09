package app

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	st "github.com/kormiltsev/tbot-kid/internal/storage"
)

var s = make(chan string)
var sleepers = 0

func Sleepy(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	if sleepers == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "When exactly?\nFormat: HHMM, like 1145 (means 11:45)")
		if a, err := bot.Send(msg); err == nil {
			return a.MessageID
		} else {
			log.Panic(err)

		}
		return 0
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Уже спит. Проснулась? /kidwu")
		if _, err := bot.Send(msg); err == nil {
			log.Panic(err)
		}
	}
	return 0
}

func WakeUp(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	if sleepers == 1 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "When exactly?")
		if a, err := bot.Send(msg); err == nil {
			return a.MessageID
		} else {
			log.Panic(err)
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Никто не спит")
		if _, err := bot.Send(msg); err == nil {
			log.Panic(err)
		}
	}
	return 0
}

func SleepEnds(update tgbotapi.Update) {
	s <- update.Message.Text

	//close(s)
}

func SleepStarts(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
	t := time.Now()
	msgid := 0
	if len(update.Message.Text) == 4 {

		hh, err := strconv.Atoi(update.Message.Text[:2])
		if err != nil || hh >= 24 {
			return
		}
		mm, err := strconv.Atoi(update.Message.Text[2:4])
		if err != nil || mm >= 60 {
			return
		}
		current_time := time.Now()
		t = time.Date(current_time.Year(), current_time.Month(), current_time.Day(),
			hh, mm, 0, 0, time.Local)
		if current_time.Before(t) {
			log.Printf("wrong time")
			return
		}
		text := fmt.Sprintf("Света заснула в %s", t.Local())
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		updt, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
		msgid = updt.MessageID
		for _, u := range Users.SU {
			if u != update.Message.Chat.ID {
				msg := tgbotapi.NewMessage(u, text)
				log.Printf("sent")
				_, err = bot.Send(msg)
				if err != nil {
					log.Panic(err)
				}
			}
		}
	}
	sleepers++
	for {
		select {
		case wutime := <-s:
			if len(wutime) == 4 {
				t1 := time.Now()
				hh, err := strconv.Atoi(wutime[:2])
				if err != nil || hh >= 24 {
					break
				}
				mm, err := strconv.Atoi(wutime[2:4])
				if err != nil || mm >= 60 {
					break
				}
				current_time := time.Now()
				t1 = time.Date(current_time.Year(), current_time.Month(), current_time.Day(),
					hh, mm, 0, 0, time.Local)
				if current_time.Before(t1) || t1.Before(t) {
					log.Printf("wrong time")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Время раньше времени засыпания.")
					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
					break
				}

				sleepers--
				data := st.Dt{
					ChatID:   strconv.FormatInt(update.Message.Chat.ID, 10),
					Perc:     "",
					Type:     "sleeping",
					Date:     t,
					Duration: t1.Sub(t),
				}
				stat := st.AddRow(data)
				ans := fmt.Sprintf("Света проснулась.\nЗаснула в %v\nСпала %v\n%s", t, t1.Sub(t), stat)
				for _, u := range Users.SU {
					msg := tgbotapi.NewMessage(u, ans)
					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
				}
			}
		case <-time.After(time.Minute * 1):
			msg := tgbotapi.NewEditMessageText(update.Message.Chat.ID, msgid, fmt.Sprintf("Света заснула в  %s\nи спит уже %v\n", t.Local(), time.Since(t)))
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}

}
