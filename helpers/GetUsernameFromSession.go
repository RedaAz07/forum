package helpers

import (
	"forum/utils"
)

func GetUsernameFromSession(token string) string {
	var username string
	err := utils.Db.QueryRow("SELECT COALESCE(username, email) FROM users WHERE session = ?", token).Scan(&username)
	if err != nil {
		return ""
	}
	return username
}
