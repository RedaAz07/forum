package route

import (
	"fmt"
	"forum/handlers"
	"forum/helpers"
	"forum/utils"
	"net/http"
	"text/template"
)

func Route() {
	var err error

	utils.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("err parsing templates", err)
		return
	}
	http.HandleFunc("/", (handlers.HomeHandler))
	http.HandleFunc("/logout", helpers.Auth(handlers.LogOutHandler))

	http.HandleFunc("/login", handlers.LoginShowHandler)
	http.HandleFunc("/register", handlers.RegisterShowHandler)

	http.HandleFunc("/loginAuth", handlers.LoginHandler)
	http.HandleFunc("/registerAuth", handlers.RegisterHandler)

	http.HandleFunc("/static/", handlers.StyleHandler)

	http.HandleFunc("/createPost", helpers.Auth(handlers.CreatePost))

	http.HandleFunc("/reaction", helpers.Auth(handlers.ReactionHandler))

	http.HandleFunc("/comment", helpers.Auth(handlers.CommentHandler))


	http.HandleFunc("/CommentsLike", helpers.Auth(handlers.CommentsLikeHandler))


	




}
