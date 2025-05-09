package bot 

import (
	"fmt"
	"os"
	"log"
	"os/signal"
	"strings"
	"BGCE-BOT/bot/commands"

	
	"github.com/bwmarrin/discordgo"
)


var Token string 

func Run() {
	Token = os.Getenv("Token")
	if Token == "" {
		log.Fatal("No Token found in the environment variable")
	}
	
	dg, err := discordgo.New("Bot " + Token)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	if err != nil {
		log.Fatal("err")
	}

	dg.AddHandler(newMessage)

	dg.Open()
	defer dg.Close()

	fmt.Println("Bot Running...")
	c := make(chan os.Signal, 1) 
	signal.Notify(c, os.Interrupt)
	<-c 
}

func newMessage(discordSession *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == discordSession.State.User.ID {
		return 
	}

	
	content := strings.ToLower(message.Content)


	switch {
	case strings.Contains(message.Content, "!bot"):
		discordSession.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hi there! (Responded in %v)"))
	
	case strings.HasPrefix(content, "!meme"):
		commands.SendMeme(discordSession, message.ChannelID) 

	case strings.HasPrefix(content, "!ask"):
		query := strings.TrimSpace(strings.TrimPrefix(content, "!ask"))
		commands.Ask(discordSession, message, query)

	}
		
}


