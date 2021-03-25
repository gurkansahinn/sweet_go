package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func MessageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print(m.Content)
}
