package main

import (
	"fmt"
	route "forum/routes"
	"net/http"
)

func main() {
fmt.Println("Starting server on port 8080")
route.Route()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
		return

	}

}
