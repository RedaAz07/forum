package handlers

import (
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusNotFound, utils.ErrorNotFound)
		return
	}
	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		sessValue = ""
	} else {
		sessValue = session.Value
	}
	query := `select id ,  session from users where session = ?`
	var userId int
	sess := ""
	utils.Db.QueryRow(query, sessValue).Scan(&userId, &sess)

	sessValue = sess
	// get comments
	commentMap, err := helpers.FetchComments(r)
	categorMap, errcat := helpers.FetchCategories()
	categories, errall := helpers.AllCategories()
	if err != nil || errcat != nil || errall != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	// !  get posts
	stmt := `SELECT 
				p.id, 
				p.username, 
				p.title, 
				p.description, 
				p.time, 
				COALESCE((p.image_path), '') AS imaghe_path,
				COUNT(CASE WHEN l.value = 1 THEN 1 END) AS total_likes, 
				COUNT(CASE WHEN l.value = -1 THEN 1 END) AS total_dislikes,
				COALESCE((
					SELECT value FROM likes WHERE postID = p.id AND userID = ?
				), 0) AS user_reaction_pub
				FROM posts p
				LEFT JOIN likes l ON p.id = l.postID
				GROUP BY p.id
				ORDER BY p.time DESC;
	`
	rows, err := utils.Db.Query(stmt, userId)
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	var posts []utils.Posts
	var post utils.Posts
	var totalLikes, totalDislikes, user_reaction_pub int
	for rows.Next() {

		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &post.ImagePath, &totalLikes, &totalDislikes, &user_reaction_pub)
		if err != nil {
			helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
			return
		}
		post.Categories = categorMap[post.Id]
		post.Comments = commentMap[post.Id]
		post.TotalLikes = totalLikes
		post.TotalDislikes = totalDislikes
		post.TotalComments = len(commentMap[post.Id])
		post.UserReactionPosts = user_reaction_pub

		now := time.Now()
		diff := now.Sub(post.Time)
		seconds := int(diff.Seconds())
		post.TimeFormatted = helpers.FormatDuration((seconds))
		posts = append(posts, post)

	}

	// !  end get posts

	variables := struct {
		Session    string
		UserActive string
		Posts      []utils.Posts
		Categories []utils.Categories
	}{
		Session:    sessValue,
		UserActive: helpers.GetUsernameFromSession(sessValue),
		Posts:      posts,
		Categories: categories,
	}

	helpers.RanderTemplate(w, "home.html", http.StatusOK, variables)
}
