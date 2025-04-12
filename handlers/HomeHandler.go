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
                    COUNT(CASE WHEN l.value = -1 THEN 1 END) AS total_dislikes, 
                    COALESCE(c.comment, '') AS comment,
                    COALESCE(c.username, '') AS username,
                    COALESCE(c.time, '') AS time
             FROM posts p
             LEFT JOIN likes l ON p.id = l.postID
             LEFT JOIN comments c ON p.id = c.postID
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
		var comment utils.Comments
		
		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &totalLikes, &totalDislikes, &comment.Comment, &comment.Username, &comment.Time)
		if err != nil {
			fmt.Println("Scan error:", err)
			helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
			return
		}
		
		post.Comments = append(post.Comments, comment)
		post.TotalLikes = totalLikes
		post.TotalDislikes = totalDislikes
		posts = append(posts, post)
	}

	// طباعة عدد المنشورات
	fmt.Println("Total posts:", len(posts))

	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}

	variables := struct {
		Session string
		Posts   []utils.Posts
	}{
		Session: sessValue,
		Posts:   posts,
	}

	// رندرة القالب
	helpers.RanderTemplate(w, "home.html", 200, variables)
}
