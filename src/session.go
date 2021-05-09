package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const Version = "v0.0.0-alpha"

var Session, _ = discordgo.New()

var DBumpTimeout *time.Timer

func StartSession() {
	Session.Token = "Bot " + os.Getenv("TOKEN")

	fmt.Printf(`Sweet-GO %-16s`+"\n\n", Version)

	err := Session.Open()
	if err != nil {
		log.Fatalf("Discord ile bağlantı kurulamadı, %s\n", err)
	}

	Session.AddHandler(MessageEvent)

	DBumpTimeout = time.AfterFunc(121*time.Second*1000, BumpNotify)
}
