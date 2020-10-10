package main

import (
	"github.com/bwmarrin/discordgo"
)

var commands []Command
var (conf Config)

// Stores the Bot Token
type Config struct {
	BotToken string `json:"bot_token"`
}
// Discord Message Structure Handling
type Command struct {
	Names       []string
	Description string
	Syntax      string

	Do func(ctx Context)
}
type Context struct {
	Command Command

	Session *discordgo.Session
	Guild   string
	Author  *discordgo.User
	Channel string
	Args    []string
}

func init() {
	commands = append(commands, Command{Names: []string{"snipe"}, Syntax: "snipe", Description: "Select a Username to Claim", Do: execTurbo})
	commands = append(commands, Command{Names: []string{"search"}, Syntax: "search", Description: "Searches NameMC for 10 Upcoming Names", Do: execSearch})
}
