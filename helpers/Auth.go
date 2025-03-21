package helpers

import (
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value == "" {
			RanderTemplate(w, "login.html", http.StatusUnauthorized, nil)
			return
		}

		HandlerFunc(w, r)
	}
}
