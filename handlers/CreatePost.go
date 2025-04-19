package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// !  check the post method
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, nil)
		return

	}
	// !  check the session that if the user is logged in
	exists, session := helpers.SessionChecked(w, r)
	if !exists {
		fmt.Println("session exists", exists)
		http.Redirect(w, r, "/", 302)
		return
	}
	// !  get the data

	err := r.ParseForm()
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return

	}
	title := r.FormValue("title")
	description := r.FormValue("description")

	category := r.Form["tags"] //* if he just choose the category

	fmt.Println("category", category)
	// !  check if the user create a new category







	// ! get the username
	stmt2 := `select  username from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username)

	stmt := `insert into posts (title, description, username) values(?, ?, ?)`

	utils.Db.Exec(stmt, title, description, username)

	http.Redirect(w, r, "/", 302)
}
