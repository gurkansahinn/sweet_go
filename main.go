package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	discord "github.com/lilAmper/sweet-go/src"
)

const Version = "v0.0.0-alpha"

var Session, _ = discordgo.New()

func main() {
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf(".env dosyası yüklenemedi.")
		return
	}

	Session.Token = os.Getenv("TOKEN")

	fmt.Printf(`Sweet-GO%-16s`+"\n\n", Version)

	flag.Parse()

	if Session.Token == "" {
		log.Println("Discord bot tokeni boş girilmiş gözüküyor.")
		return
	}

	err = Session.Open()
	if err != nil {
		log.Printf("Discord ile bağlantı kurulamadı, %s\n", err)
		os.Exit(1)
	}

	Session.AddHandler(discord.MessageEvent)

	log.Printf(`Sweet-GO çalıştırıldı. CTRL + C ile kapatabilirsin.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}
