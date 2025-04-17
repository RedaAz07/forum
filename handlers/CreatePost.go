package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
	title := r.FormValue("title")
	description := r.FormValue("description")
	category := r.FormValue("categoryID")//* if he just choose the category
	createCategory := r.FormValue("categories")// * if he create a new category
// !  check if the user create a new category

if createCategory != ""  && category == "" {
	//  insert the new category into the database
	stmt := `insert into categories (name) values(?)`
	_, err := utils.Db.Exec(stmt, createCategory)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}
	// get the id of the new category
	stmt2 := `select id from categories where name = ?`
	row := utils.Db.QueryRow(stmt2, createCategory)
	err = row.Scan(&category)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	 
}


	categoryID, _ := strconv.Atoi(category)

	// ! get the username 
	stmt2 := `select  username from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username)

	stmt := `insert into posts (title, description, categoryID, username) values(?, ?, ?, ?)`

	_, err := utils.Db.Exec(stmt, title, description, categoryID, username)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return

	}
	http.Redirect(w, r, "/", 302)
}