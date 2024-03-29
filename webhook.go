package logdis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendLogToDiscord(whurl string, dw *discordWebhook) {
	j, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
}

func setWebhookStruct(name string, img string) *discordWebhook {
	dw := &discordWebhook{
		UserName:  name,
		AvatarURL: img,
	}
	return dw
}

func setWebfookMessage(dis *discordWebhook, message string, user string, level string) *discordWebhook {
	dis.Content = user
	dis.Embeds = []discordEmbed{
		discordEmbed{
			Title: level,
			Desc:  message,
			Color: 0x550000,
		},
	}
	return dis
}

// Required to use discord webhook with Go
type discordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}
type discordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}
type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type discordEmbed struct {
	Title  string         `json:"title"`
	Desc   string         `json:"description"`
	URL    string         `json:"url"`
	Color  int            `json:"color"`
	Image  discordImg     `json:"image"`
	Thum   discordImg     `json:"thumbnail"`
	Author discordAuthor  `json:"author"`
	Fields []discordField `json:"fields"`
}

type discordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []discordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
}
