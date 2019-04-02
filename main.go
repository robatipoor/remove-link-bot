package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

var port string
var token string
var appURL string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	appURL = os.Getenv("APP_URL")
	if appURL == "" {
		log.Fatal("Application URL not set")
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatalln("TELEGRAM_TOKEN not set !")
	}
}

func main() {

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: appURL},
	}
	setings := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	bot, err := tb.NewBot(setings)
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle(tb.OnDocument, func(m *tb.Message) {
		log.Println("Get Document Message !")
		doc := m.Document
		doc.Caption = removeAddresses(m.Caption)
		bot.Send(m.Sender, doc)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		log.Println("Get Text Message!")
		if m.Text == "/start" {
			bot.Send(m.Sender, "please send me post !")
			return
		}
		caption := removeAddresses(m.Text)
		bot.Send(m.Sender, caption)
	})

	bot.Handle(tb.OnPhoto, func(m *tb.Message) {
		log.Println("Get Photo Message !")
		ph := m.Photo
		ph.Caption = removeAddresses(m.Caption)
		bot.Send(m.Sender, ph)
	})

	bot.Handle(tb.OnVideo, func(m *tb.Message) {
		log.Println("Get Video Message !")
		vi := m.Video
		vi.Caption = removeAddresses(m.Caption)
		bot.Send(m.Sender, vi)
	})

	bot.Handle(tb.OnAudio, func(m *tb.Message) {
		log.Println("Get Audio Message !")
		au := m.Audio
		au.Caption = removeAddresses(m.Caption)
		bot.Send(m.Sender, au)
	})

	log.Println("Start Bot ...")
	bot.Start()
}
