package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/7uu13/forum/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	models.PerformMigrations(db)


	newUser := models.User{
		Username:  "john_doe",
		Email:     "john@example.com",
		Password:  "securepassword",
		Age:       65,
		Gender:    "penguin",
		FirstName: "john",
		LastName:  "doe",
	}

	newPost := models.Post{
		Title:   "suema",
		Content: "asedfasdfasdf",
		UserId:  2,
	}

	newPost2 := models.Post{
		Title:   "kelleema",
		Content: "sdfghdfghdfg",
		UserId:  2,
	}
	if err := insertUser(db, newUser); err != nil {
		panic(err)
	}

	if err := insertPost(db, newPost); err != nil {
		panic(err)
	}

	if err := insertPost(db, newPost2); err != nil {
		panic(err)
	}

	_, posts, _ := getUsersWithPosts(db, 2)
	fmt.Println(posts)

	//user, err := getUserById(db, 1)
	//fmt.Printf("%T\n", user)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getUserById(db, 1, w, r)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}



func insertUser(db *sql.DB, user models.User) error {
	insertStmt := `INSERT INTO users (username, email, password, age, firstname, lastname, gender) VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := db.Exec(insertStmt, user.Username, user.Email, user.Password, user.Age, user.FirstName, user.LastName, user.Gender)
	return err
}

func insertPost(db *sql.DB, post models.Post) error {
	stmt := `INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)`

	_, err := db.Exec(stmt, post.Title, post.Content, post.UserId)
	return err
}

func getUserById(db *sql.DB, id int, w http.ResponseWriter, r *http.Request) {
	var user models.User
	stmt := `SELECT username, email FROM users WHERE id = ?`

	err := db.QueryRow(stmt, id).Scan(&user.Username, &user.Email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, user.Username)
}

func getUsersWithPosts(db *sql.DB, userId int) (models.User, []models.Post, error) {
	var user models.User
	var posts []models.Post

	userQuery := `SELECT id, username FROM users WHERE id = ?`
	err := db.QueryRow(userQuery, userId).Scan(&user.Id, &user.Username)
	if err != nil {
		return user, nil, err
	}

	postQuery := `SELECT id, title, content FROM posts WHERE user_id = ?`
	rows, err := db.Query(postQuery, userId)
	if err != nil {
		return user, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			return user, nil, err
		}
		posts = append(posts, post)
	}

	return user, posts, nil
}
