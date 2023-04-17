package logdis

import (
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

// This structure defines what is needed to output logs to any channel on discord.
type Logger struct {
	level Level
	mutex sync.Mutex

	// This is the webhook url of the channel for which you want to send notifications to the discord.
	// ex) discord.com/api/webhooks/xxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	Webhook string

	// This is the url of the icon image of the bot that sends notifications to the discord
	// ex) https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikiniku0115/20190806/20190806000644.png
	Img string

	// This is the name of the bot that will send notifications to the discord
	// ex) hogehoge
	Name string
}

func NewLogger(webhook string, img string, name string) *Logger {
	return &Logger{
		level:   InfoLevel,
		Webhook: webhook,
		Img:     img,
		Name:    name,
	}
}

func (l *Logger) check(level Level) bool {
	return l.Level() >= level
}

func (l *Logger) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

type Fields map[string]string

func (l *Logger) Log(level Level, user string, args ...interface{}) {
	if l.check(level) {
		message := ""
		message = fmt.Sprint(args...)

		l.mutex.Lock()
		defer l.mutex.Unlock()
		log.Println(message)

		// send log to discord
		dis := setWebhookStruct(l.Name, l.Img)
		dis = setWebfookMessage(dis, message, user, level.String())
		sendLogToDiscord(l.Webhook, dis)
	}
}

func (l *Logger) Logf(level Level, user string, format string, args ...interface{}) {
	if l.check(level) {
		message := ""
		message = fmt.Sprintf(format, args...)

		l.mutex.Lock()
		defer l.mutex.Unlock()
		log.Println(message)

		// send log to discord
		dis := setWebhookStruct(l.Name, l.Img)
		dis = setWebfookMessage(dis, message, user, level.String())
		sendLogToDiscord(l.Webhook, dis)
	}
}

func (l *Logger) Info(user string, i ...interface{}) {
	l.Log(InfoLevel, user, i...)
}

func (l *Logger) Infof(user string, format string, i ...interface{}) {
	l.Logf(InfoLevel, user, format, i...)
}

func (l *Logger) Debug(user string, i ...interface{}) {
	l.Log(DebugLevel, user, i...)
}

func (l *Logger) Debugf(user string, format string, i ...interface{}) {
	l.Logf(DebugLevel, user, format, i...)
}

func (l *Logger) Error(user string, i ...interface{}) {
	l.Log(ErrorLevel, user, i...)
}

func (l *Logger) Errorf(user string, format string, i ...interface{}) {
	l.Logf(ErrorLevel, user, format, i...)
}

func (l *Logger) Warn(user string, i ...interface{}) {
	l.Log(WarnLevel, user, i...)
}

func (l *Logger) Warnf(user string, format string, i ...interface{}) {
	l.Logf(WarnLevel, user, format, i...)
}

func (l *Logger) Fatal(user string, i ...interface{}) {
	l.Log(FatalLevel, user, i...)
	os.Exit(1)
}

func (l *Logger) Fatalf(user string, format string, i ...interface{}) {
	l.Logf(FatalLevel, user, format, i...)
	os.Exit(1)
}

func (l *Logger) Panic(user string, i ...interface{}) {
	l.Log(PanicLevel, user, i...)
	panic(fmt.Sprint(i...))
}

func (l *Logger) Panicf(user string, format string, i ...interface{}) {
	l.Logf(PanicLevel, user, format, i...)
	panic(fmt.Sprintf(format, i...))
}
