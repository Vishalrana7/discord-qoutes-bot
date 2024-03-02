package bot

import (
	"discord-qoute-bot/config"
	"fmt"
	"math/rand"
	"time"
	"github.com/bwmarrin/discordgo"
	
)

var BotID string
var goBot *discordgo.Session

// Array to store qoutes

var qoutes = []string{
	"Never outshine the master",
	"Never put too much trust in friends: learn how to use enemies",
	"Conceal your intentions",
	"Always say less than necessary",
	"So much depends on reputation - Guard it with your life",
	"Court attention at all costs",
	"Get Others to do the work for you, but always take the credit",
	"Make other people come to you use bait if necessary",
	"Win through your actions, Never through argument",
	"Infection: avoid the unhappy amd unlucky",
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		// _, _ = s.ChannelMessageSend(m.ChannelID, "pong")

		//random qoutes
		rand.Seed(time.Now().Unix())
		randomIndex := rand.Intn(len(qoutes))
		randomQoute := qoutes[randomIndex]


		_, _ = s.ChannelMessageSend(m.ChannelID, randomQoute)


	}
}