package handlers

import (
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	_, sees := helpers.SessionChecked(w, r)

	req := utils.Db.QueryRow(`select id from users  where   session =  ?`, sees)
	userid := ""
	req.Scan(&userid)

	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, nil)
		return
	}
	postID := r.FormValue("postID")
	if postID == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	query := ` DELETE FROM posts WHERE id = ?   and userID =  ? `
	res, err := utils.Db.Exec(query, postID, userid)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
