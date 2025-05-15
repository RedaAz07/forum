package helpers

import (
	"net/http"

	"forum/utils"
)

func AllCategories(w http.ResponseWriter) []utils.Categories {
	var categories []utils.Categories

	stmt := `SELECT name, id FROM categories`
	rows, err := utils.Db.Query(stmt)
	if err != nil {
		RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var category utils.Categories
		err := rows.Scan(&category.Name, &category.Id)
		if err != nil {
			RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return nil
		}
		categories = append(categories, category)
	}

	return categories
}
