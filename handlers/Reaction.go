package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, session := helpers.SessionChecked(w, r)

	

	fmt.Println("session exists", session)

	postID := r.FormValue("postID")
	reaction := r.FormValue("reaction")

	if postID == "" || reaction == "" || (reaction != "1" && reaction != "-1") {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	stmt2 := "SELECT id  FROM users WHERE session = ?"
	var userid int
	errr := utils.Db.QueryRow(stmt2, session).Scan(&userid)

	if errr!=nil {
		fmt.Println("Error scanning reaction 111 11 1 1 1 1 :", errr)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
fmt.Println("userid", userid)
fmt.Println("postID", postID)


	stmt := `select value from likes where postID = ? and  userID = ?`
	row := utils.Db.QueryRow(stmt, postID, userid)
	var reactionval int
	err := row.Scan(&reactionval)
	if err != nil {
		



		//make the like 

		stmt := `insert into likes (postID, userID, value) values(?, ?, ?)`
		_, err := utils.Db.Exec(stmt, postID, userid, reaction)
		if err != nil {
			fmt.Println("Error inserting reaction:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
http.Redirect(w,r,"/", 302)
	}
fmt.Println("reactionval", reactionval)
http.Redirect(w,r,"/", 302)
}
