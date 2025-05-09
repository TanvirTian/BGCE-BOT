package main 

import (
	"BGCE-BOT/bot"
	"log"
	"os"

)


func main() {
	token, ok := os.LookupEnv("Token")
	if !ok || token == "" {
		log.Fatal("Must set Discord token as env variable: Token")
	}

	bot.Token = token
	bot.Run()


}
