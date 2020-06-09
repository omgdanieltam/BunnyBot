/*
Functions and struct for the authentication tokens
*/
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/buger/jsonparser"
)

type Auth struct {
	discord string
	imgur string
	wolfram string
}

// build the auth struct
func build_auth() {
	// read the auth json file
	json, err := ioutil.ReadFile("auth.json")
	if err != nil {
		fmt.Print("Error reading auth.json file, ")
		panic(err)
	}

	// grab the discord token
	discord, err := jsonparser.GetString(json, "[0]", "discord")
	if err != nil {
		fmt.Print("Error parsing discord token, ")
		panic(err)
	}
	auth.discord = discord

	// grab the imgur token
	imgur, err := jsonparser.GetString(json, "[0]", "imgur")
	if err != nil {
		fmt.Print("Error parsing imgur token, ")
		panic(err)
	}
	auth.imgur = imgur

	// grab the wolfram alpha token
	wolfram, err := jsonparser.GetString(json, "[0]", "wolfram")
	if err != nil {
		fmt.Print("Error parsing wolfram token, ")
		panic(err)
	}
	auth.wolfram = wolfram
}
