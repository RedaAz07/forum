package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// get all users
	var allusers []string
	var username string
	queryUsers := `select username from users order by username ASC`
	users, errUsers := utils.Db.Query(queryUsers)
	for users.Next() {
		errUsers = users.Scan(&username)
		if errUsers != nil {
			fmt.Println("DB Query error:", errUsers)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		allusers = append(allusers, username)
	}
	// end of get all usrers

	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		// fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}
	query := `select id from users where session = ?`
	var userId int
	utils.Db.QueryRow(query, sessValue).Scan(&userId)

	// get comments
	commentMap := helpers.FetchComments(w, r)

	categorMap := helpers.FetchCategories(w)

	//! end of the map
	// get user id to use it in commentlikes and publikes

	// get categories
	categories := helpers.AllCategories(w)
	// end get categories
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
		fmt.Println("DB Query error:", err)
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return
	}

	var posts []utils.Posts
	var post utils.Posts
	var totalLikes, totalDislikes, user_reaction_pub int
	for rows.Next() {

		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &post.ImagePath, &totalLikes, &totalDislikes, &user_reaction_pub)
		if err != nil {
			fmt.Println("Scan error:", err)
			helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
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
		PostCatgs  []string
		Users []string
	}{
		Session:    sessValue,
		UserActive: helpers.GetUsernameFromSession(sessValue),
		Posts:      posts,
		Categories: categories,
		Users: allusers,
	}

	helpers.RanderTemplate(w, "home.html", 200, variables)

	/* !  bnisba l ay relation many to many endi hna kaykhnsi njbdha bohdha 7it adir lya mochkil f posts f like w dislikes
	so ghatl9ani jbedt l categories w l comments f bohdhom w7tithom fwa7d lmap bach key hya post id w value hya struct
	mn b3d kan7thom fstruct post li fiha kolchi
	*/
	//!  i add the clike on comments + add table  +add the form in home html but its not working + handler without fixing
}
