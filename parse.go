package logdis

import (
	"fmt"
)

func parse(f Fields, message string, user string) string {
	if !CheckDefaultUser(user) {
		user = "@" + user
	}

	var resstr string
	resstr = user + "\n"

	for key, value := range f {
		str := fmt.Sprintf("%s : %s", key, value)
		resstr += str
	}
	resstr += " { message : "
	resstr += message
	resstr += "}"

	return resstr
}
