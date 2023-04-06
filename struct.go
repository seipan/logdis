package logdis

// This structure defines what is needed to output logs to any channel on discord.
type Logger struct {

	// This is the webhook url of the channel for which you want to send notifications to the discord.
	// ex) discord.com/api/webhooks/xxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	Webhook string

	// This is the url of the icon image of the bot that sends notifications to the discord
	// ex) https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikiniku0115/20190806/20190806000644.png
	Img string

	// This is the name of the bot that will send notifications to the discord
	// ex) hogehoge
	Name string
}

func NewLogger(webhook string, img string, name string) *Logger {
	return &Logger{
		Webhook: webhook,
		Img:     img,
		Name:    name,
	}
}
