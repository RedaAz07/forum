package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func CommentsLikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, session := helpers.SessionChecked(w, r)

	commentId := r.FormValue("commentID")
	userId:= r.FormValue("userId")
	reaction := r.FormValue("reaction")
	if userId == "" || commentId ==""|| reaction == "" || (reaction != "1" && reaction != "-1") {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	stmt2 := "SELECT id  FROM users WHERE session = ?"
	var userid int
	errr := utils.Db.QueryRow(stmt2, session).Scan(&userid)

	if errr != nil {
		fmt.Println("Error scanning reaction 111 11 1 1 1 1 :", errr)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}


	stmt := `select value from commentsLikes
	
	inner join comments on comments.id = commentslikes.commentID
	inner join  posts on posts.id = comments.postID
	where commentID = ? and  userID = ?`
	row := utils.Db.QueryRow(stmt, commentId, userid)
	var reactionValue string
	err := row.Scan(&reactionValue)
	if err != nil {

		// make the like

		stmt := `insert into commentsLikes (commentID, userID, value) values(?, ?, ?)`
		_, err := utils.Db.Exec(stmt, commentId, userid, reaction)
		if err != nil {
			fmt.Println("Error inserting reaction:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", 302)
	} else {
		if reactionValue == reaction {
			// delete the like
			stmt := `delete from commentsLikes where commentID = ? and userID = ?`
			_, err := utils.Db.Exec(stmt, commentId, userid)
			if err != nil {
				fmt.Println("Error deleting reaction:", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", 302)
			return
		} else {

			// update the like
			stmt := `update commentsLikes set value = ? where commentID = ? and userID = ?`
			_, err := utils.Db.Exec(stmt, reaction, commentId, userid)
			if err != nil {
				fmt.Println("Error updating reaction:", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", 302)
			return
		}
	}
}
