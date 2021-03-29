package bot

import (
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	message "github.com/lilAmper/sweet-go/src/messages"
)

func MessageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Content) == "sweet" {
		message.SweetMessage(s, m)
	} else if strings.ToLower(m.Content) == "!d bump" {
		DbumpMessage(s, m)
	} else if string(m.Content[0]) == "!wiki" {
		log.Print("COMMAND HANDLER")
	}
}

func DbumpMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer DBumpTimeout.Stop()

	DBumpTimeout = time.AfterFunc(121*time.Second, BumpNotify)
}
