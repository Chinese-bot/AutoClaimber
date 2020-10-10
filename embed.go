package main

import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)


func (ctx Context) send(variety int, channel, message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendEmbed(channel, defaultEmbed(variety).SetDescription(message).MessageEmbed)
}

func (ctx Context) reply(variety int, message string) (*discordgo.Message, error) {
	return ctx.send(variety, ctx.Channel, message)
}


func (ctx Context) showSyntax() {
	_, _ = ctx.reply(3, "Please define the Username you wish to Snipe")
}

func (ctx Context) showError() {
	_, _ = ctx.reply(3, "Error in Function")
}

func (ctx Context) showError1() {
	_, _ = ctx.reply(3, "Please define the search volume")
}

func defaultEmbed(variety int) *embed.Embed {
	e := embed.NewEmbed()

	if variety == 0 { /* Search Result */
		e.SetColor(0x8000FF).SetAuthor("Chinese Turbo Search Results").SetFooter("Chinese Turbo made by chinese#2661")
	} else if variety == 1 { /* Information */
		e.SetColor(0xF000FF).SetAuthor("Chinese Turbo Searching...").SetFooter("Chinese Turbo made by chinese#2661")
	} else if variety == 2 { /* Warning */
		e.SetColor(0xFFDE33).SetAuthor("Starting Snipe...").SetFooter("Chinese Turbo made by chinese#2661")
	} else if variety == 3 { /* Error */
		e.SetColor(0xFF3636).SetAuthor("Chinese Turbo Error").SetFooter("Chinese Turbo made by chinese#2661")
	} else if variety == 4 { /* Error */
		e.SetColor(0x8000FF).SetAuthor("Chinese Turbo Sniping...").SetFooter("Chinese Turbo made by chinese#2661")
	} else if variety == 5 { /* Error */
		e.SetColor(0x8000FF).SetAuthor("Sniping...").SetFooter("Chinese Turbo made by chinese#2661")
	}

	return e
}

