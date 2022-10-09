package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddUser(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
	args := strings.Split(update.Message.Text, "_")
	newuser, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "if you want to add user, send me ID like this:\n/adduser_233322233"))
		return
	}
	Users.US = append(Users.US, newuser)
	mss := "Actual list of IDs in group of Users\n"
	for _, uss := range Users.US {
		mss = mss + strconv.FormatInt(uss, 10) + "\n"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, mss)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func AddSuperUser(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
	args := strings.Split(update.Message.Text, "_")
	newuser, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "if you want to add user, send me ID like this:\n/adduser_233322233"))
		return
	}
	Users.SU = append(Users.SU, newuser)
	mss := "Actual list of IDs in group of SuperUsers\n"
	for _, uss := range Users.SU {
		mss = mss + strconv.FormatInt(uss, 10) + "\n"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, mss)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func ChechUser(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) string {
	u := update.Message.Chat.ID
	for _, us := range Users.SU {
		if u == us {
			return "SU"
		}
	}
	for _, us := range Users.US {
		if u == us {
			return "User"
		}
	}

	for _, u := range Users.SU {
		t := fmt.Sprintf("[!] Unknown user start this bot: %s, to add as SUPERUSER /kidaddsuperuser_%v\nto add as USER /kidadduser_%v", update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Chat.ID)
		msg := tgbotapi.NewMessage(u, t)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	return "Unknown"

}

func ListUser(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
	t := "SuperUsers:\n"
	for _, us := range Users.SU {
		t = t + fmt.Sprintf("%v\n", us)
	}
	t = t + fmt.Sprintf("users:\n")
	for _, us := range Users.US {
		t = t + fmt.Sprintf("%v\n", us)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, t)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
