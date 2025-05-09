package main 

import (
	"BGCE-BOT/bot"
	"log"
	"os"

)


func main() {
	Token := os.Getenv("Token")
	if Token == "" {
		log.Fatal("Must set Discord token as env variable: Token")
	}

	bot.Token = Token
	bot.Run()
}
