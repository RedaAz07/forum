package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	stmt := `SELECT p.id, p.username, p.title, p.description, p.time, 
                    COUNT(CASE WHEN l.value = 1 THEN 1 END) AS total_likes, 
                    COUNT(CASE WHEN l.value = -1 THEN 1 END) AS total_dislikes
             FROM posts p
             LEFT JOIN likes l ON p.id = l.postID
             GROUP BY p.id
             ORDER BY p.time DESC`

	rows, err := utils.Db.Query(stmt)
	if err != nil {
		fmt.Println("DB Query error:", err)
		helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
		return
	}

	var posts []utils.Posts

	for rows.Next() {
		var post utils.Posts
		var totalLikes, totalDislikes int
		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &totalLikes, &totalDislikes)
		if err != nil {
			fmt.Println("Scan error:", err)
			helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
			return
		}
		post.TotalLikes = totalLikes
		post.TotalDislikes = totalDislikes
		post.TimeFormatted = post.Time.Format("2006-01-02 15:04:05")
		posts = append(posts, post)
	}
	// fmt.Println("Total posts:", len(posts))

	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		// fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}

	variables := struct {
		Session  string
		Username string
		Posts    []utils.Posts
	}{
		Session:  sessValue,
		Username: helpers.GetUsernameFromSession(sessValue), // une fonction pour récupérer le nom
		Posts:    posts,
	}
	

	helpers.RanderTemplate(w, "home.html", 200, variables)
}
