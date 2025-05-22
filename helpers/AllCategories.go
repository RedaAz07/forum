package helpers

import (
	"forum/utils"
)

func AllCategories() ([]utils.Categories, error) {
	var categories []utils.Categories

	stmt := `SELECT name, id ,icon FROM categories`
	rows, err := utils.Db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category utils.Categories
		err := rows.Scan(&category.Name, &category.Id, &category.Icon)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
