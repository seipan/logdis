package logdis

type Fields map[string]interface{}

func Info(f Fields, message string, user string) {
	str := parse(f, message, user)

}

func Infoln(Fields, message string, user string) {

}

func Debug(Fields, message string, user string) {

}

func Debugln(Fields, message string, user string) {

}

func Error(Fields, message string, user string) {

}

func Errorln(Fields, message string, user string) {

}

func Warn(Fields, message string, user string) {

}

func Warnln(Fields, message string, user string) {

}

func Fatal(Fields, message string, user string) {

}

func Fatalln(Fields, message string, user string) {

}

func Panic(Fields, message string, user string) {

}

func Panicln(Fields, message string, user string) {

}
