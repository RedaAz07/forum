package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func Filter_By_Categorie(w http.ResponseWriter, r *http.Request) {
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

	// !  end

	session, errr := r.Cookie("session")
	var sessValue string
	if errr != nil {
		// fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}

	err := r.ParseForm()
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return

	}
	categories := r.Form["tags"]
	if len(categories) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	query := `  SELECT  
				p.id, 
				p.username, 
				p.title, 
				p.description, 
				p.time
				FROM posts p
				INNER JOIN categories_post cp ON p.id = cp.postID
				INNER JOIN categories c ON cp.categoryID = c.id
				WHERE c.id = ?
				ORDER BY p.time DESC
`
	var posts []utils.Posts
	var post utils.Posts
	mapp := make(map[int]bool) // pour éviter duplication

	for _, categorie := range categories {
		rows, err := utils.Db.Query(query, categorie)
		if err != nil {
			fmt.Println("query error", err)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		for rows.Next() {
			err := rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time)
			if err != nil {
				fmt.Println(" error", err)
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
				return
			}
			if !mapp[post.Id] {
				post.Categories = categorMap[post.Id]
				posts = append(posts, post)
				mapp[post.Id] = true
				now := time.Now()
				diff := now.Sub(post.Time)
				seconds := int(diff.Seconds())
				post.TimeFormatted = helpers.FormatDuration((seconds))
			}
		}
	}
	// ! get categories
	var categoriess []utils.Categories

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
		categoriess = append(categoriess, category)

	}
	//! end get categories

	variables := struct {
		Session    string
		UserActive string
		Posts      []utils.Posts
		Categories []utils.Categories
		PostCatgs  []string
	}{
		Session:    sessValue,
		UserActive: helpers.GetUsernameFromSession(sessValue),
		Posts:      posts,
		Categories: categoriess,
	}

	helpers.RanderTemplate(w, "home.html", 200, variables)
}
