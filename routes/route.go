package route

import (
	"fmt"
	"forum/handlers"
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
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/statuspage", handlers.StatusPage)

}
