package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	auth Auth
)

func init() {
	// build the authentication struct
	buildAuth()
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

	// make sure we match our bot's message token
	if len(m.Content) > 2 {
		input_token := m.Content[0:2]
		if strings.ToLower(string(input_token)) != "t." {
			return
		}
	}

	// get our message without our bot's token
	message := m.Content[2:]

	// awwnime (test)
	if message == "awwnime" {
		// get url
		url := <-get_imgur_image("awwnime")

		// print message with url
		s.ChannelMessageSend(m.ChannelID, url)
	}

	//get_redditbooru("awwnime")

}
