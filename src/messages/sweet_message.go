package message

import (
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	closestmatch "github.com/schollz/closestmatch"
)

var messages = []string{
	`Kahretsin, tekrar başlıyoruz.`,
	`Grove Street Aileleri artık büyük değil.`,
	`Grove Street, ev. En azından her şeyi berbat etmeden önce öyleydi.`,
	`Grove Street kraldır! Benimle söyleyin zenciler, Grove Street kraldır! EVET!`,
	`Gidelim.`,
	`Yapma! Ne yapıyorsun? Carl, Brian, kesin şunu!`,
	`Yanlış evi seçtin, pislik!`,
	`Kahrolası treni takip et CJ!`,
	`Sakinleş Big Smoke, SAKİNLEŞ!`,
	`Ryder haklı. Hepimiz ayrılalım ve daha sonra tekrar buluşalım.`,
	`Adamım, kimse mahalleye önem vermiyor.`,
	`Bilmiyorum adamım!`,
	`_Megafon Sesi: LSPD! aracını durdur! HEY HEY, NE YAPIYORSUN? BİZİ ÖLDÜRECEKSİN._`,
	`Hey, hey! Big Smoke, it's me, Carl! Chill! Chill!`,
	`It's time for smoke.`,
	`Mahalleni sevmen lazım, tıpkı kardeşlerini sevdiğin gibi.`,
	`Naber, kanka?`,
	`Geç kaldık, hadi koyulalım şu işe!`,
}

func SweetMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	bag := []int{len(messages)}
	cm := closestmatch.New(messages, bag)
	match := cm.Closest(strings.ToLower(m.Content))

	if sort.SearchStrings(messages, match) != 0 {
		s.ChannelMessageSend(m.ChannelID, "<#786000773497749524>, sözlerimi tekrarlama göt!")
		return
	}

	rand.Seed(time.Now().Unix())

	message := messages[rand.Intn(len(messages))]
	s.ChannelMessageSend(m.ChannelID, message)
}
