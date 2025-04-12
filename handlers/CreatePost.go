package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/helpers"
	"forum/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, nil)
		return

	}

	exists, session := helpers.SessionChecked(w, r)
	if !exists {
		fmt.Println("session exists", exists)
		http.Redirect(w, r, "/", 302)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	category := r.FormValue("categoryID")
fmt.Println("description =>" , description)

	categoryID, _ := strconv.Atoi(category)

	stmt2 := `select  username from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username)

	fmt.Println("username", username)
	stmt := `insert into posts (title, description, categoryID, username) values(?, ?, ?, ?)`

	_, err := utils.Db.Exec(stmt, title, description, categoryID, username)
	if err != nil {
		helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return

	}
	http.Redirect(w, r, "/", 302)
}




