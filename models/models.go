package models

type User struct {
	Id        int
	Username  string
	Age       int
	Gender    string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Post struct {
	Id      int
	Title   string
	Content string
	UserId  int
}
