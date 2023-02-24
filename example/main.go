package main

import (
	logdis "github.com/seipan/logdis"
)

const (
	YOURTOKEN = "xxxxxxxxxxxx"
)

func main() {
	bot := logdis.NewDiscordBot(YOURTOKEN)
	logdis.ConnectBot(bot)

	bot.Infoln(nil, "hi my first message", logdis.AllUser)
}
