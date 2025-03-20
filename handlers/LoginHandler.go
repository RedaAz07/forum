package handlers

import (
	"database/sql"
	"forum/helpers"
	"forum/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		helpers.RanderTemplate(w, "StatusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}

	username := r.FormValue("username")
	passowrd := r.FormValue("password")

	if username == "" || passowrd == "" {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, " empy data")
		return	
	}
	stmt := `SELECT password FROM users WHERE  username  =  ? OR   email  = ? `
	row := utils.Db.QueryRow(stmt, username, username)
	var hashPass string
	err := row.Scan(&hashPass)

	if err == sql.ErrNoRows {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, "invalide username ")
		return	
	}
	Err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(passowrd))
	if Err != nil {
		helpers.RanderTemplate(w, "login.html", http.StatusBadRequest, "invalide password ")
		return	
	} else {

		helpers.RanderTemplate(w, "home.html", http.StatusOK, nil)
		return

	}
	

}
