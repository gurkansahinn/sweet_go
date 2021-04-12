package bot

import (
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	message "github.com/lilAmper/sweet-go/src/messages"
	closestmatch "github.com/schollz/closestmatch"
)

func MessageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Content) == "sweet" {
		message.SweetMessage(s, m)
	} else if strings.ToLower(m.Content) == "!d bump" {
		DbumpMessage(s, m)
	} else if strings.ToLower(m.Content) == "!twitch" {
		s.ChannelMessageSend(m.ChannelID, "https://twitch.tv/amperc")
	} else {
		bag := []int{len(message.DontAskMessages)}
		cm := closestmatch.New(message.DontAskMessages, bag)
		match := cm.Closest(strings.ToLower(m.Content))

		if sort.SearchStrings(message.DontAskMessages, match) != 0 {
			s.ChannelMessageSend(m.ChannelID, "Soruyu sormak için sormamalısın.\nhttps://dontasktoask.com/tr/")
		}

		bag = []int{len(message.ReadRulesMessages)}
		cm = closestmatch.New(message.ReadRulesMessages, bag)
		match = cm.Closest(strings.ToLower(m.Content))

		if sort.SearchStrings(message.ReadRulesMessages, match) != 0 {
			s.ChannelMessageSend(m.ChannelID, "<#786000773497749524> kanalını okumadan mesaj göndermemelisin. Satış işlemlerini bu discord üzerinde yapamazsın.")
		}
	}
}

func DbumpMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer DBumpTimeout.Stop()

	DBumpTimeout = time.AfterFunc(120*time.Second, BumpNotify)
}
