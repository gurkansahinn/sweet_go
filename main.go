package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	bot "github.com/lilAmper/sweet-go/src"
)

const Version = "v0.0.0-alpha"

var Session, _ = discordgo.New()

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(".env dosyası yüklenemedi.")
		return
	}

	Session.Token = "Bot " + os.Getenv("TOKEN")

	fmt.Printf(`Sweet-GO%-16s`+"\n\n", Version)

	err = Session.Open()
	if err != nil {
		log.Printf("Discord ile bağlantı kurulamadı, %s\n", err)
		os.Exit(1)
	}

	Session.AddHandler(bot.MessageEvent)

	log.Printf(`Sweet-GO çalıştırıldı.`)
	select {}
}
