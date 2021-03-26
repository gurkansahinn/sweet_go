package main

import (
	"log"

	"github.com/joho/godotenv"
	bot "github.com/lilAmper/sweet-go/src"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(".env dosyası yüklenemedi.")
		return
	}

	bot.StartSession()
	log.Printf(`Sweet-GO çalıştırıldı.`)
	select {}
}
