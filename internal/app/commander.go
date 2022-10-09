package app

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Chats = make(map[int64]chat)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	// meet the user (who it is)

	userstatus := ChechUser(bot, update, Users)

	if userstatus == "SU" {

		// if button, add here

		// check for command
		if update.Message.IsCommand() {
			Chats[update.Message.Chat.ID] = chat{
				Ms:     0,
				Status: "none",
			}
			args := strings.Split(update.Message.Text, "_")
			switch args[0] {
			case "/kidsleep":
				gg := Sleepy(bot, update)
				Chats[update.Message.Chat.ID] = chat{
					Ms:     gg,
					Status: "sleepstarttime",
				}
				return
			case "/kidwu":
				gg := WakeUp(bot, update)
				Chats[update.Message.Chat.ID] = chat{
					Ms:     gg,
					Status: "sleependtime",
				}
				return
			case "/kidaddsuperuser":
				AddSuperUser(bot, update, Users)
				return
			case "/kidadduser":
				AddUser(bot, update, Users)
				return
			case "/kidlistofusers":
				ListUser(bot, update, Users)
				return
			case "/kidaskcatalog":
				GetRows(bot, update)
				return
			default:
				Menu(bot, update, Users)
				return
			}
		}

		st := Chats[update.Message.Chat.ID]
		switch st.Status {
		case "sleepstarttime":
			st.Status = "none"
			Chats[update.Message.Chat.ID] = st
			go SleepStarts(bot, update, Users)
			return
		case "sleependtime":
			st.Status = "none"
			Chats[update.Message.Chat.ID] = st
			SleepEnds(update)
			return
		default:

		}
		Menu(bot, update, Users)
		return
	}

	if userstatus == "User" {
		return
	}

	if update.Message == nil {
		return
	}

}
