package handlers

import (
	"database/sql"
	"forum/helpers"
	"forum/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "StatusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
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
		helpers.RanderTemplate(w, "login.html", http.StatusInternalServerError, "Database error")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password)) != nil {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, "Invalid username or password")
		return
	}

	sessionID := uuid.New().String()

	stmt2 := `UPDATE users SET session = ? WHERE username = ?`
	_, err = utils.Db.Exec(stmt2, sessionID, username)
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

	helpers.RanderTemplate(w, "home.html", http.StatusOK, nil)
}
