package logdis

type Fields map[string]interface{}

func (b *Client) Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)

}

func (b *Client) Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Client) Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)

}

func (b *Client) Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Client) Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)

}

func (b *Client) Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Client) Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)

}

func (b *Client) Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Client) Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)

}

func (b *Client) Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	str += "\n"

}

func (b *Client) Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)

}

func (b *Client) Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	str += "\n"

}
