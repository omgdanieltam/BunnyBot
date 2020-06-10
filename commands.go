/*
This contains all of our commands used throughout the bot except for image gathering
*/

package main

import (
	"math/rand"
)

// flip a coin
func coinflip(author string) string{
	// get a random 1 or 0
	num := rand.Intn(1)

	// return string based on number
	if num == 0 { // heads
		return "<@" + author + "> flipped a coin, it landed on **heads!**"
	} else { // tails
		return "<@" + author + "> flipped a coin, it landed on **tails!**"
	}
}