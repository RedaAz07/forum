package main

import (
	"fmt"
	"html/template"
	"net/http"

	"forum/helpers"
	route "forum/routes"
	"forum/utils"
)

func main() {
	helpers.DataBase()
	var err error

	utils.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("err parsing templates", err)
		return
	}
	fmt.Println("server listening on http://localhost:8080/")
	route.Route()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
}
