package handlers

import (
	"forum/helpers"
	"forum/utils"
	"net/http"
	"time"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
helpers.RanderTemplate(w,"login.html", 200, nil)
		return
	}

	_, err = utils.Db.Exec("Update users set session = ? where session = ?", "Null", cookie.Value)
	if err != nil {
		helpers.RanderTemplate(w,"home.html", http.StatusInternalServerError, nil)
		return		
	}

	expiredCookie := http.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &expiredCookie)

helpers.RanderTemplate(w,"login.html",200, nil )
}
