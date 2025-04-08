package helpers


func GetAllPosts() ([][]string, error) {
	rows, err := Db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}
