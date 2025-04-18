package helpers

import (
	"fmt"

	"forum/utils"
)

func GetUsernameFromSession(token string) string {
	var username string
	err := utils.Db.QueryRow("SELECT COALESCE(username, email) FROM users WHERE session = ?", token).Scan(&username)
	if err != nil {
		fmt.Println("GetUsernameFromSession error:", err)
		return ""
	}
	return username
}
