package handlers

import (
	"database/sql"
	"forum/helpers"
	"forum/utils"
	"net/http"
	"regexp"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

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

	stmt := "SELECT  id FROM users where username = ?  "

	row := utils.Db.QueryRow(stmt, username)
	var id string
	err := row.Scan(&id)

	if err != sql.ErrNoRows {
		ErrorMessage = "the username is already used  "

	}

	

	if ErrorMessage != "" {

		helpers.RanderTemplate(w, "login.html", http.StatusBadGateway, ErrorMessage)
		return
	}

	stmt2 := `INSERT INTO users (username, email, password) VALUES (?, ?, ?);`

	_, err = utils.Db.Exec(stmt2, username, email, passowrd)
	if err != nil {
		helpers.RanderTemplate(w, "login.html", 200, " try again ")
		return
	}

	helpers.RanderTemplate(w, "home.html", 200, nil)

}
