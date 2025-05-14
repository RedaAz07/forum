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
		fmt.Println(photoURL)
	} else {
		fmt.Println("ddddd")
		photoURL = ""
	}
	///////////////
	if file!=nil{
			buffer := make([]byte, 512)

_,_ = file.Read(buffer)
	if err != nil {
		http.Error(w, "Can't read file", http.StatusInternalServerError)
		return
	}

	file.Seek(0, io.SeekStart)

	contentType := http.DetectContentType(buffer)

	if !strings.HasPrefix(contentType, "image/")  {
		fmt.Println(contentType)
		http.Redirect(w, r, "/", 302)

		return

	}
	}
	////////////////////////
	
	

	//! end of upload

	///////////////////////////////////////////
	title := r.FormValue("title")
	description := r.FormValue("description")

	category := r.Form["tags"] //* if he just choose the category


	stmt2 := `select  username from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username)

	stmt := `insert into posts (title, description, username,image_path) values(?, ?, ?, ?)`

	res, _ := utils.Db.Exec(stmt, title, description, username, photoURL)

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
