package logdis

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

func CheckDefaultUser(user User) bool {
	if user == DebugUser || user == AdminUser || user == DevUser || user == AllUser || user == HereUser {
		return true
	}

	return false
}
