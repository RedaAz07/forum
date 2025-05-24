package handlers

import (
	"fmt"
	"forum/helpers"
	"forum/utils"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, nil)
		return
	}
	postID := r.FormValue("postID")
	if postID == "" {
		fmt.Println("333f")
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, "Missing post ID")
		return
	}

	query := `DELETE FROM posts WHERE id = ?`
	_, err := utils.Db.Exec(query, postID)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
	}

	http.Redirect(w, r, "/", 302)
}

// {{if eq $.UserActive .Username}} <!-- Only show to post owner -->
// <form action="/deletePost" method="post">
//   <input type="hidden" name="postID" value="{{.Id}}" />
//   <button type="submit" class="delete-btn">
// 	<i class="fa-solid fa-trash-can"></i>
//   </button>
// </form>
// {{end}}
