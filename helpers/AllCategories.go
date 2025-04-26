package helpers

import (
	"net/http"

	"forum/utils"
)

func AllCategories(w http.ResponseWriter) []utils.Categories {
	// ! get categories
	var categories []utils.Categories

	stmtcategpries := `SELECT name, id FROM categories `
	rows3, err3 := utils.Db.Query(stmtcategpries)
	if err3 != nil {
		RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return nil

	}

	for rows3.Next() {
		var category utils.Categories
		err3 = rows3.Scan(&category.Name, &category.Id)
		if err3 != nil {
			RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return nil
		}
		categories = append(categories, category)

	}
	return categories
	// ! end get categories
}
