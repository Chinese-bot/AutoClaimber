package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Start The Bot
func init() {
	bytes, err := ioutil.ReadFile("config.json")

	if err != nil {
		println("Error loading the file 'config.json'", err)
	}
	json.Unmarshal(bytes, &conf)
	if conf.BotToken == "" {
		fmt.Println("No Discord Token is Set")
		return
	}

	bytes, err = ioutil.ReadFile("names.json")
	if err != nil {
		println("Error loading the file 'names.json'", err)
	}
	json.Unmarshal(bytes, &Storage)
}

// Grab Bot Token, Storing Discord handlers // Opening Bot to WebSocket
func main() {
	var err error
	dg, err := discordgo.New("Bot " + conf.BotToken)

	if err != nil {
		fmt.Println("Discord Bot Session was Unavailable to start", err)
		return
	}
	// message handler function
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot || !strings.HasPrefix(m.Message.Content, ">") {
			return
		}
		args := strings.Split(m.Message.Content[1:], " ")
		for _, c := range commands {
			for _, n := range c.Names {
				if strings.ToLower(args[0]) == n {
					c.Do(Context{Session: s, Author: m.Author, Command: c, Channel: m.ChannelID, Guild: m.GuildID, Args: args[1:]})
					return
				}
			}
		}
	})

	ready := make(chan bool, 1)
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		ready <- true
	})

	if err = dg.Open(); err != nil {
		panic(err)
	}

	<-ready

	_, _ = fmt.Scanln()
	_ = dg.Close()
}

var Storage = make(map[string]time.Time)


func execTurbo(ctx Context) {
	if len(ctx.Args) != 1 {
		ctx.showSyntax()
		return
	}
	_, _ = ctx.reply(2, "🎉 "+fmt.Sprint(ctx.Args)+" is releasing @ "+Storage[strings.ToLower(ctx.Args[0])].Format("Mon, 02 Jan 2006 15:04:05"))
	// need to store the accounts to snipe with and check if the account name is changable
	// check the accounts every 10minutes to ensure that they can be used

	//prepare accounts for snipe upon >Snipe

	// put the username in queue

	//snipe the username
}

func execSearch(ctx Context) {
	if len(ctx.Args) != 1 {
		ctx.showError1()
		return
	}

	_, _ = ctx.reply(1, "NameMC Usernames with Monthly/Searches > "+fmt.Sprint(ctx.Args)+"")
	proxy := Premium()
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}, Timeout: time.Second * 25}
	req, _ := http.NewRequest("GET", "https://namemc.com/minecraft-names?length_op=&length=&lang=&searches="+ctx.Args[0], nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		ctx.showError()
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200{
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}


	var name []string
	doc.Find("body > main > div > div.col-lg-7.order-lg-2 > div > div.card-body.p-0").Find("div").Each(func(i int, s *goquery.Selection){
		html, _ := s.Html()
		split := strings.Split(html, `href="/name/`)
		if len(split) == 1 || strings.Contains(html, "div") {
			return
		}
		name = append(name, strings.ToLower(strings.Split(split[1], `"`)[0]))
	})


	var dates []string
	doc.Find("body > main > div > div.col-lg-7.order-lg-2 > div > div.card-body.p-0").Find("div").Each(func(i int, s *goquery.Selection){
		html, _ := s.Html()
		split := strings.Split(html, `datetime="`)
		if len(split) == 1 || strings.Contains(html, "div") {
			return
		}
		dates = append(dates, strings.Split(split[1], `"`)[0])

	})


	var summary string
	for i, name := range name {
		t, _ := time.Parse(time.RFC3339, dates[i])
		t = t.Add(time.Hour)
		summary += name + " - " + t.Format("02/01/2006, 15:04:05") + "\n"
		Storage[name] = t

		//Storage.Names = append(Storage.Names, Name{name, t}

	}
	file, _ := json.Marshal(Storage)
	_ = ioutil.WriteFile("names.json", file, 0644)
	ctx.reply(0, summary)

	//time.Sleep(time.Now().Sub(Releases) - time.Millisecond * 100)
}
































/*
message, err := ctx.Session.ChannelMessageSendComplex(session.Channel, &discordgo.MessageSend{
        Embed:   defaultEmbed(1).SetDescription("React below with the type of social media platform deal you wish to proceed.").SetTitle("Establish the Deal").MessageEmbed,
        Content: "<@" + session.Participants[0] + "> + <@" + session.Participants[1] + ">",
    })
 */






