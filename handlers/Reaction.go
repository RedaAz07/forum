package handlers

import "net/http"

func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	reactionType := r.FormValue("reaction_type")


	
	if postID == "" || reactionType == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reaction recorded successfully"))
}
