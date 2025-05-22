package handlers

import (
	"database/sql"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}
	_, session := helpers.SessionChecked(w, r)

	postID := r.FormValue("postID")
	reaction := r.FormValue("reaction")
	if postID == "" || reaction == "" || (reaction != "1" && reaction != "-1") {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return
	}

	stmt2 := "SELECT id  FROM users WHERE session = ?"
	var userid int
	errr := utils.Db.QueryRow(stmt2, session).Scan(&userid)

	if errr != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)

		return
	}

	stmt := `select value from likes where postID = ? and  userID = ?`
	row := utils.Db.QueryRow(stmt, postID, userid)
	var reactionValue string
	err := row.Scan(&reactionValue)
	if err == sql.ErrNoRows {
		// make the like
		stmt := `insert into likes (postID, userID, value) values(?, ?, ?)`
		_, err := utils.Db.Exec(stmt, postID, userid, reaction)
		if err != nil {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
			return
		}
		http.Redirect(w, r, "/", 302)
	} else {
		if reactionValue == reaction {
			// delete the like
			stmt := `delete from likes where postID = ? and userID = ?`
			_, err := utils.Db.Exec(stmt, postID, userid)
			if err != nil {
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)

				return
			}
			http.Redirect(w, r, "/", 302)
			return
		} else {
			// update the like
			stmt := `update likes set value = ? where postID = ? and userID = ?`
			_, err := utils.Db.Exec(stmt, reaction, postID, userid)
			if err != nil {
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
				return
			}

			http.Redirect(w, r, "/", 302)
			return
		}
	}
}
