package handlers

import (
	"database/sql"
	"forum/helpers"
	"forum/utils"
	"net/http"
	"regexp"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "StatusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}

	password := r.FormValue("password")
	email := r.FormValue("email")
	username := r.FormValue("username")
	firstpassword := r.FormValue("firstpassword")

	var ErrorMessage string

	emailregex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	//passregex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`

	if password == "" || email == "" || username == "" || firstpassword == "" {
		ErrorMessage = "All inputs are required"
	} else if match, _ := regexp.MatchString(emailregex, email); !match {
		ErrorMessage = "Invalid email format"
	} else if firstpassword != password {
		ErrorMessage = "Passwords do not match"
	}  else if len(username) < 8 {
		ErrorMessage = "Username must be at least 8 characters"
	}

	stmt := "SELECT id FROM users WHERE username = ? OR email = ?"
	row := utils.Db.QueryRow(stmt, username, email)
	var id string
	err := row.Scan(&id)

	if err != sql.ErrNoRows {
		ErrorMessage = "The username or email is already used"
		helpers.RanderTemplate(w, "register.html", http.StatusBadRequest, ErrorMessage)
		return
	}

	if ErrorMessage != "" {
		helpers.RanderTemplate(w, "register.html", http.StatusBadRequest, ErrorMessage)
		return
	}

	hashPassword, Err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if Err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	stmt2 := `INSERT INTO users (username, email, password) VALUES (?, ?, ?);`
	_, err = utils.Db.Exec(stmt2, username, email, string(hashPassword))
	if err != nil {
		helpers.RanderTemplate(w, "register.html", http.StatusBadRequest, "Try again")
		return
	}

	// إنشاء السيشن مباشرة بعد التسجيل
	sessionID := uuid.New().String()
	stmt3 := `UPDATE users SET session = ? WHERE username = ?`
	_, err = utils.Db.Exec(stmt3, sessionID, username)
	if err != nil {
		helpers.RanderTemplate(w, "register.html", http.StatusInternalServerError, "Error creating session. Please try again later.")
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
