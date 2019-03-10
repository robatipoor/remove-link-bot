package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func handleBot() *tb.Bot {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
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
	return bot
}
