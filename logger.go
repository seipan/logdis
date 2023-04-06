package logdis

import "log"

type Fields map[string]interface{}

func (b *Logger) Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	log.Print(str)
}

func (b *Logger) Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	str += "\n"
	log.Print(str)

}

func (b *Logger) Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	log.Print(str)

}

func (b *Logger) Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	str += "\n"
	log.Print(str)

}

func (b *Logger) Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	log.Print(str)

}

func (b *Logger) Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	str += "\n"
	log.Print(str)

}

func (b *Logger) Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	log.Print(str)

}

func (b *Logger) Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	str += "\n"
	log.Print(str)

}

func (b *Logger) Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	log.Print(str)

}

func (b *Logger) Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	str += "\n"
	log.Print(str)

}

func (b *Logger) Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	log.Print(str)

}

func (b *Logger) Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	str += "\n"
	log.Print(str)

}
