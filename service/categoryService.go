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

func GetCategoryBySlug(db *sql.DB, slug string) ([]model.Categories, error) {
	// Comment
	var categories []model.Categories
	var category model.Categories
	stmt := `SELECT * FROM categories WHERE name_slug = ?`

	err := db.QueryRow(stmt, slug).Scan(&category.Id, &category.Name, &category.Name_slug)
	categories = append(categories, category)

	if err != nil {
		if err == sql.ErrNoRows {
			return categories, err
		}
		return categories, err
	}
	return categories, nil
}
