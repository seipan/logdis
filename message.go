package logdis

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientId := os.Getenv("CLIENT_ID")
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	if u.ID != clientId {
		b.sendMessage(s, m.ChannelID, "hello logger")
	}
}

func (b *Bot) sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
