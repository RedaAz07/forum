package middleware

import (
	"net/http"

	"forum/utils"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			// Check if the session is valid
			stmt := "SELECT id FROM users WHERE session = ?"
			var userID int
			err = utils.Db.QueryRow(stmt, cookie.Value).Scan(&userID)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			
		}

		HandlerFunc(w, r)
	}
}
