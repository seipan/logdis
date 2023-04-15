package logdis

import (
	"fmt"
)

func parse(f Fields, message string, user User) string {
	if !CheckDefaultUser(user) {
		user = "@" + user
	}

	var resstr string
	resstr = user.String() + " "

	for key, value := range f {
		str := fmt.Sprintf("{ %s : %s }", key, value)
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
	dis.Content = user
	disf := []discordField{}
	for i, v := range f {
		d := discordField{
			Name:  i,
			Value: v,
		}
		disf = append(disf, d)
	}

	dis.Embeds = []discordEmbed{
		discordEmbed{
			Title:  level,
			Desc:   message,
			Color:  0x550000,
			Fields: disf,
		},
	}
	return dis
}
