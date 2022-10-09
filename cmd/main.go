package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kormiltsev/tbot-kid/internal/app"
)

func main() {
	godotenv.Load()
	myToken := os.Getenv("TOKEN")
	MYID := os.Getenv("MYID") // i am a superuser by default
	if myToken == "" {
		log.Printf("input telegram bot unique token:")
		fmt.Fscan(os.Stdin, &myToken)
	}
	if MYID == "" {
		log.Printf("input mine telegram bot unique ID: [none]")
		fmt.Fscan(os.Stdin, &MYID)
	}
	if MYID == "none" {
		log.Printf("run bot with SuperUsers from memory (if exists)")
	}
	s, err := strconv.ParseInt(MYID, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	Users := app.GetSuperUsers(s)

	bot, err := tgbotapi.NewBotAPI(myToken)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	// update.Message.Chat.UserName
	// update.Message.Chat.ID
	// update.Message.Text
	// update.Message.Photo
	//tgbotapi.update.Message.Chat.UserName
	//inputMessage.Chat.UserName
	//update.CallbackQuery.Data
	//update.CallbackQuery.Message.Chat.ID

	//update.CallbackQuery.Message.Message.ID
	//update.CallbackQuery.Message.MessageID

	for update := range updates {
		if update.CallbackQuery == nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		app.HandleUpdate(bot, update, Users)
	}
}
