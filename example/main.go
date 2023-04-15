package main

import "github.com/seipan/logdis"

const (
	YOURWEBHOOKURL = "xxxxxxxxxxxx"
)

func main() {
	log := logdis.NewLogger(YOURWEBHOOKURL, "", "logForExample")

	// output
	// INFO :  map[message:this is info log for test times:1]test message
	log.Info(logdis.AllUser.String(), "INFO :  ", map[string]string{"message": "this is info log for test", "times": "1"}, "test message ")

	// output
	// hoge ore
	log.Infof(logdis.AllUser.String(), "hoge%s", " ore")
}
