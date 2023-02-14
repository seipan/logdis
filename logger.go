package logdis

type Fields map[string]interface{}

func (b *Bot) Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)

	b.sendMessage(b.bot, "", str)
}

func (b *Bot) Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Bot) Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)

}

func (b *Bot) Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Bot) Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)

}

func (b *Bot) Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	str += "\n"
}

func (b *Bot) Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)

}

func (b *Bot) Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	str += "\n"
}

func (b *Bot) Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)

}

func (b *Bot) Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	str += "\n"
}

func (b *Bot) Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)

}

func (b *Bot) Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	str += "\n"
}
