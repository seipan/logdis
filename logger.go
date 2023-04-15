package logdis

import "log"

type Fields map[string]string

func (b *Logger) Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, ParseUser(user))
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "INFO")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, ParseUser(user))
	str += "\n"
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "INFO")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, ParseUser(user))
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "DEBUG")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, ParseUser(user))
	str += "\n"
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "DEBUG")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, ParseUser(user))
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "ERROR")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, ParseUser(user))
	str += "\n"
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "ERROR")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, ParseUser(user))
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "WARN")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, ParseUser(user))
	str += "\n"
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "WARN")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, ParseUser(user))
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "FATAL")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, ParseUser(user))
	str += "\n"
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "FATAL")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, ParseUser(user))
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "PANIC")
	sendLogToDiscord(b.Webhook, dis)
}

func (b *Logger) Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, ParseUser(user))
	str += "\n"
	log.Print(str)

	dis := setWebhookStruct(b.Name, b.Img)
	dis = setWebfookMessage(dis, f, message, user, "PANIC")
	sendLogToDiscord(b.Webhook, dis)
}
