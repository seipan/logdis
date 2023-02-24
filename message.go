package logdis

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientId := os.Getenv("CLIENT_ID")
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	b.channelId = m.ChannelID
	if u.ID != clientId {
		b.sendMessage(s, m.ChannelID, "hello logger")
	}
}

func (b *Bot) sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		panic(fmt.Errorf("Error sending message: %w", err))
	}
}
