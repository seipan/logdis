package logdis

import "fmt"

type User string

const (
	DebugUser User = "@Debug"
	AdminUser User = "@Admin"
	DevUser   User = "@Dev"
	AllUser   User = "@everyone"
	HereUser  User = "@here"
)

func (u User) String() string {
	return string(u)
}

func ParseUser(user string) User {
	return User(user)
}

func ToDiscordUserFormat(user string) string {
	return "@" + user
}

func CheckDefaultUser(user User) bool {
	if user == DebugUser || user == AdminUser || user == DevUser || user == AllUser || user == HereUser {
		return true
	}

	return false
}

func NewUser(user string) (User, error) {
	ok := CheckDefaultUser(ParseUser(user))
	if !ok {
		return "", fmt.Errorf("this user already used")
	}
	user = ToDiscordUserFormat(user)
	return ParseUser(user), nil
}
