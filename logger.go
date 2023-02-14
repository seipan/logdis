package logdis

type Fields map[string]interface{}

func (*Bot) Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)

}

func (*Bot) Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	str += "\n"

}

func (*Bot) Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)

}

func (*Bot) Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	str += "\n"

}

func (*Bot) Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)

}

func (*Bot) Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	str += "\n"
}

func (*Bot) Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)

}

func (*Bot) Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	str += "\n"
}

func (*Bot) Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)

}

func (*Bot) Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	str += "\n"
}

func (*Bot) Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)

}

func (*Bot) Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	str += "\n"
}
