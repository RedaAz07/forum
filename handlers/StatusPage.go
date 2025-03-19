package handlers

import (
	"forum/utils"
	"net/http"
)

func StatusPage(w http.ResponseWriter, r *http.Request) {
	err:=utils.Tp.ExecuteTemplate(w, "statuspage.html", nil)
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	
}
}
