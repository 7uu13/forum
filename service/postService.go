package service

import (
	"database/sql"
	"fmt"

	"github.com/7uu13/forum/model"
)

func CreatePost(db *sql.DB, post model.Post) (int64, error) {
	insertStmt := `INSERT INTO posts (title, content, created, user_id) VALUES (?, ?, ?, ?)`

	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(post.Title, post.Content, post.Created, post.UserId)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func GetCategoryPosts(db *sql.DB, category model.Categories) ([]model.Post, error) {
	stmt := `
	SELECT * FROM posts_category
	INNER JOIN posts ON posts_categories.post_id = posts.id
	`
	var posts []model.Post
	res, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}

	defer res.Close()

	for res.Next() {
		var post model.Post
		err = res.Scan(&post.Id, &post.Title, &post.Content, &post.Created, &post.UserId)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	fmt.Println(posts)
	return posts, nil
}
