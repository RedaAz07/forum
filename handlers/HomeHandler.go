package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	//!  get comments
	stmtcommnts := `
	SELECT 
		COALESCE(comments.comment, '') AS comment,
		COALESCE(comments.time, '') AS time,
		COALESCE(comments.username, '') AS username,
		comments.postID
	FROM comments
	INNER JOIN posts ON comments.postID = posts.id;
	`

	rows2, err2 := utils.Db.Query(stmtcommnts)
	if err2 != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return
	}

	var comments []utils.Comments

	for rows2.Next() {
		var comment utils.Comments
		err2 = rows2.Scan(&comment.Comment, &comment.Time, &comment.Username, &comment.PostID)
		if err2 != nil {
			fmt.Println("Scan error:", err2)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		comments = append(comments, comment)
	}
//  !  end get comments
// !  add the communts to   map 

commentMap := make(map[int][]utils.Comments)
	for _, c := range comments {
		commentMap[c.PostID] = append(commentMap[c.PostID], c)
	}

	// !  get posts
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
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
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

		post.Comments = commentMap[post.Id]
		post.TotalLikes = totalLikes
		post.TotalDislikes = totalDislikes
		post.TimeFormatted = post.Time.Format("2006-01-02 15:04:05")
		posts = append(posts, post)
	}


	
// !  end get posts





// ! get categories
var categories []utils.Categories

stmtcategpries := `SELECT name, id FROM categories`
	rows3, err3 := utils.Db.Query(stmtcategpries)
	if err3 != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return	
		
	}


	for rows3.Next(){
		var category utils.Categories
		err3 = rows3.Scan(&category.Name, &category.Id)
		if err3 != nil {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		categories = append(categories, category)
		
	}



	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		// fmt.Println("Session cookie error:", err)
		fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}

	variables := struct {
		Session  string
		Username string
		Posts    []utils.Posts
		Categories []utils.Categories
	}{
		Session:  sessValue,
		Username: helpers.GetUsernameFromSession(sessValue), // une fonction pour récupérer le nom
		Posts:    posts,
		Categories: categories,
	}
	

	helpers.RanderTemplate(w, "home.html", 200, variables)
}
