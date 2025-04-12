package helpers

import (
	"forum/utils"
	"net/http"
)

func SessionChecked(w http.ResponseWriter, r *http.Request) (bool, string ) {
	sessionCookie, err := r.Cookie("session")
	if err != nil || sessionCookie.Value == "" {
		return false , ""
	}

	var userID int
	stmt := "SELECT id FROM users WHERE session = ?"
	err = utils.Db.QueryRow(stmt, sessionCookie.Value).Scan(&userID)
	if err != nil {
		return false, ""
	}
	return true, sessionCookie.Value
}
