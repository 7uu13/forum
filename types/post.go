package types

import (
	"fmt"
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
	stmt := `
	SELECT * FROM posts_category
	INNER JOIN posts ON posts_categories.post_id = posts.id
	`
	var posts []Post
	res, err := config.DB.Query(stmt)
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

	fmt.Println(posts)
	return posts, nil
}
