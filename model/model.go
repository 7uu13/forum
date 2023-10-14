package model

import "time"

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Post struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	UserId  int
}

type Categories struct {
	Id        int
	Name      string
	Name_slug string
}

type PostCategories struct {
	Id         int
	PostId     int
	CategoryId int
}