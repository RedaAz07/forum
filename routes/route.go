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
	http.HandleFunc("/", helpers.Auth(handlers.HomeHandler))
	http.HandleFunc("/logout", helpers.Auth(handlers.LogOutHandler))

	http.HandleFunc("/login", handlers.LoginShowHandler)
	http.HandleFunc("/register", handlers.RegisterShowHandler)

	http.HandleFunc("/Loginreq", handlers.LoginHandler)
	http.HandleFunc("/Registerreq", handlers.RegisterHandler)

	http.HandleFunc("/static/", handlers.StyleHandler)

}
