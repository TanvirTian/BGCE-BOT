package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Ask(discordSession *discordgo.Session, m *discordgo.MessageCreate, query string) {
	queryURL := "https://api.duckduckgo.com/?q=" + url.QueryEscape(query) + "&format=json"

	resp, err := http.Get(queryURL)
	if err != nil {
		discordSession.ChannelMessageSend(m.ChannelID, "Failed to reach DuckDuckGo.")
		return
	}
	defer resp.Body.Close()

	var result struct {
		Heading       string `json:"Heading"`
		AbstractText  string `json:"AbstractText"`
		AbstractURL   string `json:"AbstractURL"`
		Answer        string `json:"Answer"`
		Definition    string `json:"Definition"`
		DefinitionURL string `json:"DefinitionURL"`
		RelatedTopics []struct {
			Text     string `json:"Text"`
			FirstURL string `json:"FirstURL"`
		} `json:"RelatedTopics"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		discordSession.ChannelMessageSend(m.ChannelID, "Failed to parse response.")
		return
	}

	if result.AbstractText != "" {
		answer := fmt.Sprintf("**%s**\n%s\nMore info: %s", result.Heading, result.AbstractText, result.AbstractURL)
		discordSession.ChannelMessageSend(m.ChannelID, answer)
		return
	}

	if result.Answer != "" {
		discordSession.ChannelMessageSend(m.ChannelID, "Answer: "+result.Answer)
		return
	}

	if result.Definition != "" {
		msg := fmt.Sprintf("Definition: %s\nMore info: %s", result.Definition, result.DefinitionURL)
		discordSession.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	if len(result.RelatedTopics) > 0 {
		
		summaryURL := "https://en.wikipedia.org/api/rest_v1/page/summary/" + strings.ReplaceAll(query, " ", "_")

		resp, err := http.Get(summaryURL)
		if err != nil {
			discordSession.ChannelMessageSend(m.ChannelID, "Failed to fetch Wikipedia summary.")
			return
		}
		defer resp.Body.Close()

		var wikiResult struct {
			Title       string `json:"title"`
			Extract     string `json:"extract"`
			ContentURLs struct {
				Desktop struct {
					Page string `json:"page"`
				} `json:"desktop"`
			} `json:"content_urls"`
		}

		err = json.NewDecoder(resp.Body).Decode(&wikiResult)
		if err != nil || wikiResult.Extract == "" {
			discordSession.ChannelMessageSend(m.ChannelID, "Wikipedia summary not available.Try being more specific")
			return
		}

		msg := fmt.Sprintf("**%s**\n%s\nMore info: %s", wikiResult.Title, wikiResult.Extract, wikiResult.ContentURLs.Desktop.Page)
		discordSession.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	discordSession.ChannelMessageSend(m.ChannelID, "No results found.")
}
