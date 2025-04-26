package handlers

import (
	"fmt"
	"log"
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
	fmt.Println(category)

	// !  check if the user create a new category

	// ! get the username
	stmt2 := `select  username from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username)

	stmt := `insert into posts (title, description, username) values(?, ?, ?)`

	res, _ := utils.Db.Exec(stmt, title, description, username)

	postID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	//! create categories

	stmtcat := `insert into categories_post (categoryID,postID) values (?,?)`

	for _, v := range category {
		_, err := utils.Db.Exec(stmtcat, v, postID)
		if err != nil {
			fmt.Println("error in inserting category", err)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
			return
		}

	}

	http.Redirect(w, r, "/", 302)
}
