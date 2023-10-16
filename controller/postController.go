package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/7uu13/forum/types"
)

type PostController struct{}

var post types.Post

func (_ *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Session Authentication for this handler
	// TODO: Check if user is logged in, if not redirect to login page
	// TODO: Get User ID from session

	/*
		Post controller

		GET - Return HTML where user can create post
		GET ?category=test - Return posts from category

		DELETE - Delete post
		POST - Create post
		PUT - Update post

	*/
	const temp_user_id = 123


	categories, err := category.GetCategories()
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			Categories      []types.Categories
			CurrentCategory types.Categories
		}{
			Categories:      categories,
			CurrentCategory: category,
		}

	switch r.Method {
	case "GET":

		tmpl, err := template.ParseGlob("ui/templates/createPost.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Fatal(err)
		}

	case "POST":
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		post := &types.Post{
			Title:   title,
			Content: content,
			UserId:  temp_user_id,
		}
		

		_, err = post.CreatePost(*post)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Post created successfully!"))
	}
}
