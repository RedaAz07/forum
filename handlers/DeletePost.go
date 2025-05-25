package handlers

import (
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorInternalServerErr)
		return
	}
	postID := r.FormValue("postID")
	if postID == "" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, "Missing post ID")
		return
	}

	query := `DELETE FROM posts WHERE id = ?`
	_, err := utils.Db.Exec(query, postID)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
