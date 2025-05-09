package commands

import (
    "fmt"
    "strings"

    "github.com/bwmarrin/discordgo"
    "github.com/go-resty/resty/v2"
)

type MemeResponse struct {
    PostLink  string `json:"postLink"`
    Title     string `json:"title"`
    URL       string `json:"url"`
    Subreddit string `json:"subreddit"`
}

func HandleMessage(discordSession *discordgo.Session, message *discordgo.MessageCreate) {
    if message.Author.ID == discordSession.State.User.ID {
        return
    }

    content := strings.ToLower(message.Content)
    //fmt.Printf("Message received: %s\n", content)
 
    switch {
    case strings.Contains(content, "!bot"):
        discordSession.ChannelMessageSend(message.ChannelID, "Hi there!")
        
    case strings.HasPrefix(content, "!meme"):
        SendMeme(discordSession, message.ChannelID)
    }
}

func SendMeme(discordSession *discordgo.Session, channelID string) {
    client := resty.New()
    resp, err := client.R().SetResult(&MemeResponse{}).Get("https://meme-api.com/gimme/")


    if err != nil {
        discordSession.ChannelMessageSend(channelID, "Failed to get meme")
        return
    }

    meme := resp.Result().(*MemeResponse)
    message := fmt.Sprintf("**%s**\n%s", meme.Title, meme.URL)
    discordSession.ChannelMessageSend(channelID, message)
}

