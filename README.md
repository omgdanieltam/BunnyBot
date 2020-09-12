# GoBunnyBot
Rewrite of my Discord bot, BunnyBot, in Go.

Previous version, written in Javascript using Node.js: https://git.dtam.pw/daniel/discord-bot-js

## To do
Restore previous functionality:
 - Allow roll ranges
 - Add 8 ball
 - Add compute
 - Add voice commands

New functionalities: 
 - ~~Add caching mechanism for images~~
 - Routely delete cache
 - Make sure same image isn't repeated for x amount of time
 - ~~Add ability to make coinflip with options to replace heads/tails~~
 - Add statistics tracking

## Building standalone
Instructions for myself on how to build this for an Alpine Linux container: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build bunnybot.go commands.go images.go auth.go
