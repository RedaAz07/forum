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
	// !  check the post method
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
	// !  get the data

	err := r.ParseForm()
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return

	}
	/////////////////////////////////////
	//! start of upload

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error uploading file:", err)
		http.Error(w, "Error uploading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create directory if not exist
	os.MkdirAll("uploads", 0o775)
	//********  ntakdo mn dak  file  ikon tswira
	buffer := make([]byte, 512)
_, err = file.Read(buffer)
if err != nil {
	http.Error(w, "Can't read file", http.StatusInternalServerError)
	return
}

  file.Seek(0, io.SeekStart)

    contentType := http.DetectContentType(buffer)
   fmt.Println(contentType)

 if !strings.HasPrefix(contentType, "image/") {
	//http.Redirect(w, r, "/", 302)
		http.Redirect(w, r, "/", 302)

	
	return
	
 }

	// Create file on server
	filename := handler.Filename
	
	
	dst, err := os.Create("uploads/" + filename)
	if err != nil {
		http.Error(w, "Cannot save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy file content
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	//! end of upload 

	///////////////////////////////////////////
	title := r.FormValue("title")
	description := r.FormValue("description")

	category := r.Form["tags"] //* if he just choose the category

	// !  check if the user create a new category

	// ! get the username
	stmt2 := `select  username from users where session = ?`
	row := utils.Db.QueryRow(stmt2, session)
	var username string
	row.Scan(&username)

	stmt := `insert into posts (title, description, username,image_path) values(?, ?, ?, ?)`

	res, _ := utils.Db.Exec(stmt, title, description, username, filename)

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
