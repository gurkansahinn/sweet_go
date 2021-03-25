package bot

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
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

func MessageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Content) == "sweet" {
		rand.Seed(time.Now().Unix())

		message := messages[rand.Intn(len(messages))]
		log.Print(s.Channel(m.ChannelID))
		s.ChannelMessageSend(m.ChannelID, message)
	}
	return
}
