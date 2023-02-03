package logger

import "fmt"

func parse(f Fields, message string, user string) string {
	var resstr string

	for key, value := range f {
		str := fmt.Sprintf("%s : %s", key, value)
		resstr += str
	}
	return resstr
}
