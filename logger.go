package logdis

type Fields map[string]interface{}

func Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)

}

func Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
}

func Debug(f Fields, message string, user string) {

}

func Debugln(f Fields, message string, user string) {

}

func Error(f Fields, message string, user string) {

}

func Errorln(f Fields, message string, user string) {

}

func Warn(f Fields, message string, user string) {

}

func Warnln(f Fields, message string, user string) {

}

func Fatal(f Fields, message string, user string) {

}

func Fatalln(f Fields, message string, user string) {

}

func Panic(f Fields, message string, user string) {

}

func Panicln(f Fields, message string, user string) {

}
