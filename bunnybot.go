package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"time"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var (
	auth Auth
	redditbooru_sources []string

	// cache settings
	cache_location string = "cache/"
	cache_time = 3 * time.Hour // 3 hours
)

func init() {
	// set the randomized seed
	rand.Seed(time.Now().UnixNano())

	// build the authentication struct
	build_auth()

	// build the redditbooru sources slice
	build_redditbooru_sources()

	// attempt to make our cached
	if _, err := os.Stat(cache_location); os.IsNotExist(err) {
		os.Mkdir(cache_location, 0777)
	}
}

func main() {
	// create Discord session
	dg, err := discordgo.New("Bot " + auth.discord)
	if err != nil {
		fmt.Println("Error creating Discord session, ", err)
		return
	}

	// register the message_create func as a callback for MessageCreate events
	dg.AddHandler(message_create)

	// open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection, ", err)
		return
	}

	// wait here until term signal
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// cleanly close
	dg.Close()

}

func message_create (s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// make sure we match our bot's message token
	if len(m.Content) > 2 {
		// get the input token
		input_token := m.Content[0:2]

		// check if we match our token
		found := false

		// try to find our token
		if strings.ToLower(string(input_token)) == "b." || strings.ToLower(string(input_token)) == "//" {
			found = true
		}

		// return if this message isn't for us
		if found == false {
			return
		}

	} else { // message too short, don't do anything
		return
	}

	// get our message's content without our bot's token
	content := strings.ToLower(m.Content[2:])

	// split our message's content into parts based on space
	message := strings.Fields(content)

	// determine our actions
	if message[0] == "coinflip" || message[0] == "coin" { // flip a coin
		s.ChannelMessageSend(m.ChannelID, coinflip(m.Author.ID, content))
	} else if message[0] == "roll" {  // roll a number
		s.ChannelMessageSend(m.ChannelID, roll(m.Author.ID))
	} else if message[0] == "source" { // print source code
		s.ChannelMessageSend(m.ChannelID, source())
	} else if message[0] == "retarded" { // retarded youtube video
		s.ChannelMessageSend(m.ChannelID, "https://youtu.be/kav7tifmyTg")
	} else if message[0] == "moon" { // wsb moon stock ticker copypasta
		s.ChannelMessageSend(m.ChannelID, moon(content)) // print moon text
	} else if len(message[0]) > 0 { // as long as there is a message, try to find a picture
		// get url
		url := <-get_image(message[0])

		// make sure we have a url returned
		if len(url) > 0 {
			s.ChannelMessageSend(m.ChannelID, url)
		} else {
			s.ChannelMessageSend(m.ChannelID, "I couldn't find that, sauce?")
		}
	}
}
