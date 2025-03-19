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
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginShowHandler)
	http.HandleFunc("/register", handlers.RegisterShowHandler)

	http.HandleFunc("/Loginreq", handlers.LoginHandler)
	http.HandleFunc("/Registerreq", handlers.RegisterHandler)

	http.HandleFunc("/statuspage", handlers.StatusPage)

}
