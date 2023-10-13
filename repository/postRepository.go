package repository

// import (
// 	"github.com/7uu13/forum/model"	
// )

// func CreatePost(post model.Post) (int64, error) {
// 	insertStmt := `INSERT INTO posts (title, content, created, user_id) VALUES (?, ?, ?, ?)`

// 	stmt, err := db.Prepare(insertStmt)
// 	if err != nil {
// 		return 0, err
// 	}

// 	result, err := stmt.Exec(post.Title, post.Content, post.Created, post.UserId)
// 	if err != nil {
// 		return 0, err
// 	}

// 	postID, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return postID, nil
// }