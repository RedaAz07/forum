package route

import (
	"fmt"
	"forum/handlers"
	"forum/utils"
	"net/http"
)

func Route() {
	utils.Tp, utils.Error = utils.Tp.ParseGlob("templates/*.html")
	if utils.Error != nil {
		fmt.Println("Error parsing templates", utils.Error)
		return
	}
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

}
