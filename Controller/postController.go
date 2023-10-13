package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"
)

func CreatePost(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Session Authentication for this handler
	// TODO: Check if user is logged in, if not redirect to login page
	// TODO: Get User ID from session
	const temp_user_id = 123123123

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "Templates/createPost.html")

	case "POST":
		fmt.Println("POST")
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		post := &model.Post{
			Title:   title,
			Content: content,
			UserId:  temp_user_id,
		}

		_, err = service.CreatePost(db, *post)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Post created successfully!"))
	}
}
