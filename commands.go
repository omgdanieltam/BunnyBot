/*
This contains all of our commands used throughout the bot except for image gathering
*/

package main

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// flip a coin
func coinflip(author string, content string) string{
	// get a random 1 or 0
	num := math.Mod(float64(rand.Intn(100)), 2) // modulo of random number between 0 and 100

	// default heads/tails text in case nothing was set from the user
	heads := "heads"
	tails := "tails"

	// try to split to see if heads/tails were set by user
	message := strings.Fields(content)

	// if the custom heads/tails was requested, set it
	if(len(message) > 2) {
		heads = message[1]
		tails = message[2]
	}

	// return string based on number
	if int(num) == 0 { // heads
		return "<@" + author + "> flipped a coin, it landed on **" + heads + "!**"
	} else { // tails
		return "<@" + author + "> flipped a coin, it landed on **" + tails + "!**"
	}
}

// roll number
func roll(author string) string{
	min := 1
	max := 100
	num := rand.Intn(max - min) + min
	return "<@" + author + "> rolls a number between 1 and 100. They roll **" + strconv.Itoa(num) + "**."
}

// source code
func source() string {
	return "BunnyBot is a Discord bot written in Go. You can view the source code here: https://git.dtam.pw/daniel/GoBunnyBot"
}