package logdis

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	bot       *discordgo.Session
	channelId string
}

func NewDiscordBot(token string) *Bot {
	bot := &Bot{}
	sesion, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	bot.bot = sesion
	return bot
}

func ConnectBot(discord *Bot) {
	discord.bot.AddHandler(discord.onMessageCreate)
	err := discord.bot.Open()
	if err != nil {
		log.Println(err)
	}

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot

	err = discord.bot.Close()
	if err != nil {
		log.Println(err)
	}
}
