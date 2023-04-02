package logdis

type Logger struct {
	Webhook   string
	Img       string
	Name      string
	ChannelId string
}

func NewClient(webhook string, channnelId string, img string, name string) *Logger {
	return &Logger{
		Webhook:   webhook,
		Img:       img,
		Name:      name,
		ChannelId: channnelId,
	}
}
