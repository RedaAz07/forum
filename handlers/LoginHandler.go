package handlers

import (
	"database/sql"
	"forum/helpers"
	"forum/utils"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}
	//  check if the user is alrady logged in 
	if exists , _ :=helpers.SessionChecked(w,r) ; exists{
		http.Redirect(w,r,"/", 303)
		return
	}
	
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, "Empty data")
		return
	}

	stmt := `SELECT password FROM users WHERE username = ? OR email = ?`
	row := utils.Db.QueryRow(stmt, username, username)

	var hashPass string
	err := row.Scan(&hashPass)
	if err == sql.ErrNoRows {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, "Invalid username or password")
		return
	} else if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError,utils.ErrorInternalServerErr)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password)) != nil {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, "Invalid username or password")
		return
	}

	sessionID := uuid.New().String()
	stmt2 := `UPDATE users SET session = ? WHERE username = ? or   email = ?`
	_, err = utils.Db.Exec(stmt2, sessionID, username, username)
	if err != nil {
		helpers.RanderTemplate(w, "login.html", http.StatusInternalServerError, "Error updating session. Please try again later.")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   3600,
	})

	http.Redirect(w, r, "/", 302)
}
