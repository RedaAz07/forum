package handlers

import (
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, nil)
		return
	}

	_, session := helpers.SessionChecked(w, r)
	


	stmt2 :=`select username from users where session = ?`
	query  := utils.Db.QueryRow(stmt2, session)
	
	var username string
	query.Scan(&username)

	postID := r.FormValue("postID")
	comment := r.FormValue("comment")

	stmt := `insert into comments (postID, comment, username ) values(?, ? ,?)`
	_, err := utils.Db.Exec(stmt, postID, comment, username)
	if err != nil {
		helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}
	http.Redirect(w, r, "/", 302)
}
