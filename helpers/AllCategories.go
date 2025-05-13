package helpers

import (
	"html/template"
	"net/http"

	"forum/utils"
)

func AllCategories(w http.ResponseWriter) []utils.Categories {
	Icons := []string{
		`<i class="fa-solid fa-medal"></i>`,
		`<i class="fa-solid fa-music"></i>`,
		`<i class="fa-solid fa-film"></i>`,
		`<i class="fa-solid fa-flask"></i>`,
		`<i class="fa-solid fa-dumbbell"></i>`,
		`<i class="fa-solid fa-microchip"></i>`,
		`<i class="fa-solid fa-scissors"></i>`,
		`<i class="fa-solid fa-landmark"></i>`,
	}

	var categories []utils.Categories

	stmt := `SELECT name, id FROM categories`
	rows, err := utils.Db.Query(stmt)
	if err != nil {
		RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
		return nil
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var category utils.Categories
		err := rows.Scan(&category.Name, &category.Id)
		if err != nil {
			RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return nil
		}
		category.Icon = template.HTML(Icons[i%len(Icons)])
		i++
		categories = append(categories, category)
	}

	return categories
}
