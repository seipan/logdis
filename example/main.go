package main

import "github.com/seipan/logdis"

const (
	YOURWEBHOOKURL = "xxxxxxxxxxxx"
)

func main() {
	log := logdis.NewLogger(YOURWEBHOOKURL, "", "logForExample")
	log.Info(map[string]string{"message": "this is info log for test", "times": "1"}, "test message ", "by me")
}
