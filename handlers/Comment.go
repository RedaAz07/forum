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

	stmt2 := `select username from users where session = ?`
	query := utils.Db.QueryRow(stmt2, session)

	var username string
	errr := query.Scan(&username)
	if errr != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return

	}

	postID := r.FormValue("postID")
	comment := r.FormValue("comment")

	stmt3 := `select id from posts where id = ?`
	query3 := utils.Db.QueryRow(stmt3, postID)

	var postID2 int
	errrr := query3.Scan(&postID2)
	if errrr != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	stmt := `insert into comments (postID, comment, username ) values(?, ? ,?)`
	_, err := utils.Db.Exec(stmt, postID, comment, username)
	if err != nil {
		helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
