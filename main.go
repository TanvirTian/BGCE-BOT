package main 

import (
	"BGCE-BOT/bot"
	"log"
	"os"

	"github.com/joho/godotenv"

)


func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading  .env file")
	}

	token, ok := os.LookupEnv("Token")
	if !ok {
		log.Fatal("Must set Discord token as env variable: Token")
	}

	bot.Token = token
	bot.Run()


}