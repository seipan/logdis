package logdis

type Fields map[string]interface{}

func Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)

}

func Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	str += "\n"

}

func Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)

}

func Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	str += "\n"

}

func Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)

}

func Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	str += "\n"
}

func Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)

}

func Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	str += "\n"
}

func Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)

}

func Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	str += "\n"
}

func Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)

}

func Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	str += "\n"
}
