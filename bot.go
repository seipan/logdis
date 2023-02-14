package logdis

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	bot *discordgo.Session
}

func NewDiscordBot(token string) *Bot {
	bot := &Bot{}
	sesion, err := discordgo.New(token)
	if err != nil {
		panic(err)
	}
	bot.bot = sesion
	return bot
}

func ConnectBot(discord *Bot) {
	discord.bot.AddHandler(onMessageCreate)
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

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientId := os.Getenv("CLIENT_ID")
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	if u.ID != clientId {
		sendMessage(s, m.ChannelID, "")
		sendReply(s, m.ChannelID, "test", m.Reference())
	}
}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

func sendReply(s *discordgo.Session, channelID string, msg string, reference *discordgo.MessageReference) {
	_, err := s.ChannelMessageSendReply(channelID, msg, reference)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
