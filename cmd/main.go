package main

import (
	"fmt"
	"forum/helpers"
	route "forum/routes"
	"net/http"
)

func main() {
	helpers.Db()
	fmt.Println("Starting server on port 8080")
	route.Route()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
		return

	}

}
