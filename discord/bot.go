package discord

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func ConnectBot() {
	discord, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		log.Println(err)
	}

	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		log.Println(err)
	}

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot

	err = discord.Close()
	if err != nil {
		log.Println(err)
	}

	return
}
