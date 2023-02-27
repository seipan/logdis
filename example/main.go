package main

import (
	"github.com/seipan/logdis"
)

const (
	YOURTOKEN = "xxxxxxxxxxxx"
)

func main() {
	bot := logdis.NewDiscordBot(YOURTOKEN)
	logdis.ConnectBot(bot)

	// If you want all members to know about the info level logs, you can do it this way
	bot.Infoln(map[string]interface{}{"type": "my first message"}, "hi my first message", logdis.AllUser)

	// If you want to let your original members know about the ERROR level log, you can do so by doing this
	bot.Errorln(map[string]interface{}{"type": "my second message"}, "hi my second message", "devmember")
}
