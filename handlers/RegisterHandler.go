package handlers

import (
	"database/sql"
	"forum/helpers"
	"forum/utils"
	"net/http"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		helpers.RanderTemplate(w, "StatusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}

	passowrd := r.FormValue("password")
	email := r.FormValue("email")
	username := r.FormValue("username")
	firstpassword := r.FormValue("firstpassword")

	var ErrorMessage string

	emailregex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	//	passregex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`

	if passowrd == "" || email == "" || username == "" || firstpassword == "" {
		ErrorMessage = "all the  inputs  is requared"

	} else if match, _ := regexp.MatchString(emailregex, email); !match {
		ErrorMessage = "invalide email"
	} else if firstpassword != passowrd {
		ErrorMessage = "the  password must be the same "
	} else if len(passowrd) < 8 {
		ErrorMessage = "invalide password "
	} else if len(username) < 8 {
		ErrorMessage = "username must be more than 8 chars  "
	}
	// ! maerftch wach  n check l user name w email fmra wla kola wa7d bohdoo

	stmt := "SELECT  id FROM users where username = ?  or email = ? "

	row := utils.Db.QueryRow(stmt, username, email)
	var id string
	err := row.Scan(&id)

	if err != sql.ErrNoRows {
		ErrorMessage = "the username  is already used  "

	}

	if ErrorMessage != "" {

		helpers.RanderTemplate(w, "register.html", http.StatusBadRequest, ErrorMessage)
		return
	}

	hashPassword, Err := bcrypt.GenerateFromPassword( []byte(passowrd) , bcrypt.DefaultCost)
	if Err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	stmt2 := `INSERT INTO users (username, email, password) VALUES (?, ?, ?);`

	_, err = utils.Db.Exec(stmt2, username, email,string(hashPassword))
	if err != nil {
		helpers.RanderTemplate(w, "register.html", http.StatusBadRequest, " try again ")
		return
	}
	helpers.RanderTemplate(w, "home.html", 200, nil)

}
