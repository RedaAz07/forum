package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func Filter_By_Categorie(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return

	}
	commentMap := helpers.FetchComments(w, r)
	allusers := helpers.AllUsers(w)
	categories := helpers.AllCategories(w)
	categorMap := helpers.FetchCategories(w)

	cookie, errr := r.Cookie("session")
	var sessValue string
	if errr != nil {
		fmt.Println("Session cookie error:", errr)
		sessValue = ""
	} else {
		sessValue = cookie.Value
	}
	categ := r.Form["tags"]
	if len(categ) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	query := `  SELECT  
				p.id, 
				p.username, 
				p.title, 
				p.description, 
				p.time,
				COALESCE((p.image_path), '') AS imaghe_path,
				COUNT(CASE WHEN l.value = 1 THEN 1 END) AS total_likes, 
				COUNT(CASE WHEN l.value = -1 THEN 1 END) AS total_dislikes
				FROM posts p
				INNER JOIN categories_post cp ON p.id = cp.postID
				INNER JOIN categories c ON cp.categoryID = c.id
				LEFT JOIN likes l ON p.id = l.postID
				WHERE c.id = ?
				GROUP BY p.id
				ORDER BY p.time DESC
`
	var posts []utils.Posts
	var post utils.Posts
	mapp1 := make(map[int]bool) // pour éviter duplication
	var totalLikes, totalDislikes int

	for _, categorie := range categ {
		rows, err := utils.Db.Query(query, categorie)
		if err != nil {
			fmt.Println("query error", err)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		for rows.Next() {
			err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &post.ImagePath, &totalLikes, &totalDislikes)
			if err != nil {
				fmt.Println(" error", err)
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
				return
			}
			if !mapp1[post.Id] {
				post.Comments = commentMap[post.Id]
				post.TotalLikes = totalLikes
				post.TotalDislikes = totalDislikes
				post.TotalComments = len(commentMap[post.Id])
				now := time.Now()
				diff := now.Sub(post.Time)
				seconds := int(diff.Seconds())
				post.TimeFormatted = helpers.FormatDuration((seconds))
				post.Categories = categorMap[post.Id]
				posts = append(posts, post)
			}
		}
	}

	variables := struct {
		Session    string
		UserActive string
		Posts      []utils.Posts
		Categories []utils.Categories
		PostCatgs  []string
		Users      []string
	}{
		Session:    sessValue,
		UserActive: helpers.GetUsernameFromSession(sessValue),
		Posts:      posts,
		Categories: categories,
		Users:      allusers,
	}

	helpers.RanderTemplate(w, "home.html", 200, variables)
}
