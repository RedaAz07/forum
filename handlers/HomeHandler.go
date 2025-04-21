package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//! get comments likes

	/* stmtgetcommentslikes := `
						COUNT(CASE WHEN cl.value = 1 THEN 1 END) AS total_likes,
	                COUNT(CASE WHEN cl.value = -1 THEN 1 END) AS total_dislikes
	from commentslikes cl
	INNER JOIN comments c ON cl.commentID = c.id
	INNER JOIN posts p ON c.postID = p.id
	GROUP BY c.id` */

	//! end get comments likes

	//! get categories
	stmtCategories := `
	SELECT C.name, C.id ,  CP.postID  FROM categories C
	INNER JOIN categories_post CP ON C.id = CP.categoryID
	INNER JOIN posts P ON CP.postID = P.id
	`

	rowcat, errcat := utils.Db.Query(stmtCategories)
	if errcat != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return
	}
	var category []utils.Categories
	for rowcat.Next() {
		var categor utils.Categories
		errcat = rowcat.Scan(&categor.Name, &categor.Id, &categor.PostID)
		if errcat != nil {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		category = append(category, categor)
	}

	// !end get categories
	// ! add the categories to the map

	categorMap := make(map[int][]utils.Categories)
	for _, d := range category {
		categorMap[d.PostID] = append(categorMap[d.PostID], d)
	}

	//! example of the map
	//!  get comments

	stmtcommnts := `
	SELECT 
    c.id,
    COALESCE(c.comment, '') AS comment,
  c.time AS time,
    COALESCE(c.username, '') AS username,
    c.postID,
    COUNT(CASE WHEN cl.value = '1' THEN 1 END) AS total_likes, 
    COUNT(CASE WHEN cl.value = '-1' THEN 1 END) AS total_dislikes
FROM comments c
INNER JOIN posts p ON c.postID = p.id
LEFT JOIN commentsLikes cl ON c.id = cl.commentID
GROUP BY c.id
ORDER BY c.time DESC;

	`

	rows2, err2 := utils.Db.Query(stmtcommnts)
	if err2 != nil {
		fmt.Println("DB Query error:", err2)
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return
	}

	var comments []utils.Comments

	for rows2.Next() {
		var comment utils.Comments
		err2 = rows2.Scan(&comment.Id, &comment.Comment, &comment.Time, &comment.Username, &comment.PostID, &comment.TotalLikes, &comment.TotalDislikes)
		if err2 != nil {
			fmt.Println("Scan error:", err2)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}

		now := time.Now()
		diff := now.Sub(comment.Time)
		seconds := int(diff.Seconds())
		comment.TimeFormattedComment = helpers.FormatDuration((seconds))

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
	var catgs []string
	for rows.Next() {
		var post utils.Posts
		var totalLikes, totalDislikes int

		err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &totalLikes, &totalDislikes)
		if err != nil {
			fmt.Println("Scan error:", err)
			helpers.RanderTemplate(w, "home.html", http.StatusInternalServerError, nil)
			return
		}
		post.Categories = categorMap[post.Id]
		post.Comments = commentMap[post.Id]
		post.TotalLikes = totalLikes
		post.TotalDislikes = totalDislikes

		now := time.Now()
		diff := now.Sub(post.Time)
		seconds := int(diff.Seconds())
		post.TimeFormatted = helpers.FormatDuration((seconds))
		posts = append(posts, post)

	}

	// !  end get posts

	// ! get categories
	var categories []utils.Categories

	stmtcategpries := `SELECT name, id FROM categories `
	rows3, err3 := utils.Db.Query(stmtcategpries)
	if err3 != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return

	}

	for rows3.Next() {
		var category utils.Categories
		err3 = rows3.Scan(&category.Name, &category.Id)
		if err3 != nil {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		categories = append(categories, category)

	}
	//! end get categories

	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		// fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}
	variables := struct {
		Session    string
		Username   string
		Posts      []utils.Posts
		Categories []utils.Categories
		PostCatgs  []string
	}{
		Session:    sessValue,
		Username:   helpers.GetUsernameFromSession(sessValue),
		Posts:      posts,
		Categories: categories,
	}
	fmt.Println("PostCatgs", catgs)

	helpers.RanderTemplate(w, "home.html", 200, variables)

	/* !  bnisba l ay relation many to many endi hna kaykhnsi njbdha bohdha 7it adir lya mochkil f posts f like w dislikes
	so ghatl9ani jbedt l categories w l comments f bohdhom w7tithom fwa7d lmap bach key hya post id w value hya struct
	mn b3d kan7thom fstruct post li fiha kolchi
	*/
	//!  i add the clike on comments + add table  +add the form in home html but its not working + handler without fixing
}
