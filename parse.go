package logdis

import (
	"fmt"
)

func parse(f Fields, message string, user string) string {
	if !CheckDefaultUser(user) {
		user = "@" + user
	}

	var resstr string
	resstr = user + "\n"

	for key, value := range f {
		str := fmt.Sprintf("%s : %s", key, value)
		resstr += str
	}
	resstr += " { message : "
	resstr += message
	resstr += "}"

	return resstr
}

func setWebhookStruct(name string, img string) *discordWebhook {
	dw := &discordWebhook{
		UserName:  name,
		AvatarURL: img,
	}
	return dw
}

func setWebfookMessage(dis *discordWebhook, f Fields, message string, user string, level string) *discordWebhook {
	dis.Embeds = []discordEmbed{
		discordEmbed{
			Title:  level,
			Desc:   message,
			Color:  0x550000,
			Author: discordAuthor{Name: user},
			Fields: []discordField{
				discordField{Name: "フィールド名1", Value: "フィールド値1", Inline: true},
				discordField{Name: "フィールド名2", Value: "フィールド値2", Inline: true},
			},
		},
	}
	return dis
}
