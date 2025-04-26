package helpers

import (
	"fmt"
	"net/http"
	"time"

	"forum/utils"
)

func FetchComments(w http.ResponseWriter , r *http.Request) map[int][]utils.Comments {
	session, err := r.Cookie("session")
	var sessValue string
	if err != nil {
		// fmt.Println("Session cookie error:", err)
		sessValue = ""
	} else {
		sessValue = session.Value
	}

	query := `select id from users where session = ?`
	var userId int
	utils.Db.QueryRow(query, sessValue).Scan(&userId)

	//!  get comments

	stmtcommnts := `
	SELECT 
    c.id,
    COALESCE(c.comment, '') AS comment,
    c.time AS time,
    COALESCE(c.username, '') AS username,
    c.postID,
    COUNT(CASE WHEN cl.value = '1' THEN 1 END) AS total_likes, 
    COUNT(CASE WHEN cl.value = '-1' THEN 1 END) AS total_dislikes,
    COALESCE((
        SELECT value 
        FROM commentsLikes 
        WHERE commentID = c.id AND userID = ?
    ), 0) AS user_reaction_comment
	FROM comments c
	LEFT JOIN commentsLikes cl ON c.id = cl.commentID
	GROUP BY c.id
	ORDER BY c.time DESC;

	`

	rows2, err2 := utils.Db.Query(stmtcommnts, userId)
	if err2 != nil {
		fmt.Println("DB Query error:", err2)
		RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return nil 
	}

	var comments []utils.Comments

	for rows2.Next() {
		var comment utils.Comments
		err2 = rows2.Scan(&comment.Id, &comment.Comment, &comment.Time, &comment.Username, &comment.PostID, &comment.TotalLikes, &comment.TotalDislikes, &comment.UserReactionComment)
		if err2 != nil {
			fmt.Println("Scan error:", err2)
			RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return nil 
		}

		now := time.Now()
		diff := now.Sub(comment.Time)
		seconds := int(diff.Seconds())
		comment.TimeFormattedComment = FormatDuration((seconds))

		comments = append(comments, comment)
	}
	//  !  end get comments

	// !  add the communts to   map
	commentMap := make(map[int][]utils.Comments)
	for _, c := range comments {
		commentMap[c.PostID] = append(commentMap[c.PostID], c)
	}

	return commentMap
}
