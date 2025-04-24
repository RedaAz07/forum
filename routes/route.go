package route

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/handlers"
	"forum/middleware"
	"forum/utils"
)

func Route() {
	var err error

	utils.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("err parsing templates", err)
		return
	}
	http.HandleFunc("/", (handlers.HomeHandler))
	http.HandleFunc("/logout", middleware.Auth(handlers.LogOutHandler))

	http.HandleFunc("/login", handlers.LoginShowHandler)
	http.HandleFunc("/register", handlers.RegisterShowHandler)

	http.HandleFunc("/loginAuth", handlers.LoginHandler)
	http.HandleFunc("/registerAuth", handlers.RegisterHandler)

	http.HandleFunc("/static/", handlers.StyleHandler)

	http.HandleFunc("/createPost", middleware.Auth(handlers.CreatePost))

	http.HandleFunc("/reaction", middleware.Auth(handlers.ReactionHandler))

	http.HandleFunc("/comment", middleware.Auth(handlers.CommentHandler))

	http.HandleFunc("/CommentsLike", middleware.Auth(handlers.CommentsLikeHandler))
}
