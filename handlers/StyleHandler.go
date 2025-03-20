package handlers

import (
	"net/http"
	"os"
	"strings"

	"forum/helpers"
	utils "forum/utils"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	File, err := os.Stat(filePath)
	if err != nil || File.IsDir() {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusNotFound, utils.ErrorNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}
