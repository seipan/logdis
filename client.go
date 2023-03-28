package logdis

type Client struct {
	Webhook   string
	ChannelId string
}

func NewClient(webhook string, channnelId string) *Client {
	return &Client{
		Webhook:   webhook,
		ChannelId: channnelId,
	}
}
