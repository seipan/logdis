package logger

const (
	DebugUser = "@Debug"
	AdminUser = "@Admin"
	DevUser   = "@Dev"
	AllUser   = "@everyone"
	HereUser  = "@here"
)

func CheckUser(user string) bool {
	if user == DebugUser || user == AdminUser || user == DevUser || user == AllUser || user == HereUser {
		return true
	}

	return false
}
