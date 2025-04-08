package helpers

import (
	"net/http"

	"forum/utils"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value == "" {
			RanderTemplate(w, "home.html", http.StatusUnauthorized, nil)
			return
		} else {

			// Check if the session is valid
			stmt := "SELECT id FROM users WHERE session = ?"
			var userID int
			err = utils.Db.QueryRow(stmt, cookie.Value).Scan(&userID)
			if err != nil {
				RanderTemplate(w, "home.html", http.StatusUnauthorized, nil)
				return
			}
		}

		HandlerFunc(w, r)
	}
}
