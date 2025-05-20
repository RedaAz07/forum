package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"forum/helpers"
	"forum/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, nil)
		return

	}
	// !  check the session that if the user is logged in
	exists, session := helpers.SessionChecked(w, r)
	if !exists {
		fmt.Println("session exists", exists)
		http.Redirect(w, r, "/", 302)
		return
	}

	err := r.ParseForm()
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return

	}

	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("myFile")
	if file != nil {
		zize := header.Size

		// fmt.Println(zize)
		maxzize := int64(10485760)
		if zize >= maxzize {
			fmt.Println("walooooo")
			helpers.RanderTemplate(w, "statusPage.html", 400, utils.ErrorBadReq)
			return

		}
	}

	var photoURL string

	if err == nil {
		defer file.Close()

		photoDir := "uploads/"

		if _, err := os.Stat(photoDir); os.IsNotExist(err) {
			err := os.MkdirAll(photoDir, 0o755)
			if err != nil {
				http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
				return
			}
		}

		photoPath := photoDir + header.Filename

		dst, err := os.Create(photoPath)
		if err != nil {
			fmt.Println("Error saving file:", err)
			http.Error(w, "Error saving photo", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		photoURL = photoDir + header.Filename
	} else {
		photoURL = ""
	}

	if photoURL != "" {
		if !strings.HasSuffix(photoURL, ".jpg") && !strings.HasSuffix(photoURL, ".png") && !strings.HasSuffix(photoURL, ".jpeg") {
			helpers.RanderTemplate(w, "statusPage.html", 400, utils.ErrorBadReq)
			return
		}
	}

	//! end of upload

	title := r.FormValue("title")
	description := r.FormValue("description")

	category := r.Form["tags"] //* if he just choose the category
	var userId int
	stmt2 := `select  username , id  from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username, &userId)

	stmt := `insert into posts (title, description, username,image_path, userID) values(?, ?, ?, ?,?)`

	res, _ := utils.Db.Exec(stmt, title, description, username, photoURL, userId)

	postID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	//! create categories

	stmtcat := `insert into categories_post (categoryID,postID) values (?,?)`

	for _, v := range category {
		_, err := utils.Db.Exec(stmtcat, v, postID)
		if err != nil {
			fmt.Println("error in inserting category", err)
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
			return
		}

	}

	http.Redirect(w, r, "/", 302)
}
