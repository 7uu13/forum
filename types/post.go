package types

import (
	"database/sql"
	"time"

	"github.com/7uu13/forum/config"
)

type Post struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	UserId  int
}

func (p *Post) CreatePost(post Post) (int64, error) {
	insertStmt := `INSERT INTO posts (title, content, created, user_id) VALUES (?, ?, ?, ?)`

	stmt, err := config.DB.Prepare(insertStmt)
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

func (p *Post) GetCategoryPosts(category Categories) ([]Post, error) {

	if category.Id == 0 {
		return nil, nil
	}

	stmt := `
	SELECT posts.*
	FROM posts
	JOIN posts_category ON posts.id = posts_category.post_id
	WHERE posts_category.category_id = ?
	`

	var posts []Post
	res, err := config.DB.Query(stmt, category.Id)
	if err != nil {
		panic(err)
	}

	defer res.Close()

	for res.Next() {
		var post Post
		err = res.Scan(&post.Id, &post.Title, &post.Content, &post.Created, &post.UserId)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Post) GetPostById(id string) (Post, error) {
	var post Post
	stmt := `SELECT * FROM posts WHERE id = ?`

	err := config.DB.QueryRow(stmt, id).Scan(&post.Id, &post.Title, &post.Content, &post.Created, &post.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return post, err
		}
		return post, err
	}
	return post, nil
}
