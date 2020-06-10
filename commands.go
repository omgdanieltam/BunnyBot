/*
This contains all of our commands used throughout the bot except for image gathering
*/

package main

import (
	"math"
	"math/rand"
	"strconv"
)

// flip a coin
func coinflip(author string) string{
	// get a random 1 or 0
	num := math.Mod(float64(rand.Intn(100)), 2) // modulo of random number between 0 and 100

	// return string based on number
	if int(num) == 0 { // heads
		return "<@" + author + "> flipped a coin, it landed on **heads!**"
	} else { // tails
		return "<@" + author + "> flipped a coin, it landed on **tails!**"
	}
}

// roll number
func roll(author string) string{
	min := 1
	max := 100
	num := rand.Intn(max - min) + min
	return "<@" + author + "> rolls a number between 1 and 100. They roll **" + strconv.Itoa(num) + "**."
}