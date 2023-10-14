package service

import (
	"database/sql"

	"github.com/7uu13/forum/model"
)

func GetCategories(db *sql.DB) ([]model.Categories, error) {
	var categories []model.Categories
	stmt := `SELECT * FROM categories`

	res, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var category model.Categories
		if err := res.Scan(&category.Id, &category.Name, &category.Name_slug); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
