package helpers

import (
	"net/http"

	"forum/utils"
)

func FetchCategories(w http.ResponseWriter) (map[int][]utils.Categories, error) {
	//! get categories
	stmtCategories := `
		SELECT C.name, C.id ,  CP.postID  FROM categories C
		INNER JOIN categories_post CP ON C.id = CP.categoryID
		INNER JOIN posts P ON CP.postID = P.id
		`

	rowcat, errcat := utils.Db.Query(stmtCategories)
	if errcat != nil {
		RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return nil, errcat
	}
	var category []utils.Categories
	for rowcat.Next() {
		var categor utils.Categories
		errcat = rowcat.Scan(&categor.Name, &categor.Id, &categor.PostID)
		if errcat != nil {
			RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
			return nil, errcat
		}
		category = append(category, categor)
	}

	// !end get categories
	// ! add the categories to the map

	categorMap := make(map[int][]utils.Categories)
	for _, d := range category {
			
			categorMap[d.PostID] = append(categorMap[d.PostID], d)
		
	}
	//! end of the map
	return categorMap, nil
}
