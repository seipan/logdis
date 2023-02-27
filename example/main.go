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

	// if you
	bot.Infoln(map[string]interface{}{"type": "my first message"}, "hi my first message", logdis.AllUser)

	bot.Errorln(map[string]interface{}{"type": "my second message"}, "hi my second message", "devmember")
}
