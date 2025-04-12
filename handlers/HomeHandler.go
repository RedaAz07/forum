package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type variable struct {
		Session string
		Posts   []utils.Posts
	}

	stmt := `select username, title, description, time from posts`
	rows, err := utils.Db.Query(stmt)
	if err != nil {
		fmt.Println("DB Query error:", err)
		helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
		return
	}
	var posts []utils.Posts

	for rows.Next() {
		var post utils.Posts
		err = rows.Scan(&post.Username, &post.Title, &post.Description, &post.Time)
		if err != nil {
			fmt.Println("Scan error:", err)
			helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
			return
		}
		posts = append(posts, post)
	}
	fmt.Println("Total posts:", len(posts))

	session, err := r.Cookie("session")

var sessValue string
if err != nil {
    fmt.Println("Session cookie error:", err)
    sessValue = "" 
} else {
    sessValue = session.Value
}

	variables := variable{
		Session: sessValue,
		Posts:   posts,
	}

	helpers.RanderTemplate(w, "home.html", 200, variables)
}
