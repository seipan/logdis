## logdis
golang logger for discord

## Usage
First, get the webhook id of the discord channel you want the service to log

![image](https://user-images.githubusercontent.com/88176012/230644885-460a7b7d-ab28-4a53-9df3-85b894527fdd.png)
The code below outputs the info level code.


You can use it by applying the url you got to YOURWEBHOOKURL on the code.

```go
package main

import "github.com/seipan/logdis"

const (
	YOURWEBHOOKURL = "xxxxxxxxxxxx"
)

func main() {
	log := logdis.NewLogger(YOURWEBHOOKURL, "", "logForExample")
	log.Info(logdis.AllUser.String(), "INFO :  ", map[string]string{"message": "this is info log for test", "times": "1"}, "test message ")
```



## Install
```go
import "github.com/seipan/logdis"
```
