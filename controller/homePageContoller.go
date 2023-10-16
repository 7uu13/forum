package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/7uu13/forum/types"
)

type HomePageController struct{}

var category types.Categories
var postRating types.PostRating

func (_ *HomePageController) HomePage(w http.ResponseWriter, r *http.Request) {
	// Define the data structure to pass to templates
	data := struct {
		Categories          []types.Categories
		CurrentCategory     types.Categories
		CurrentPost         types.Post
		CurrentPostDislikes int
		CurrentPostLikes    int

		Posts []types.Post
	}{
		Categories:          []types.Categories{},
		CurrentCategory:     types.Categories{},
		CurrentPost:         types.Post{},
		CurrentPostDislikes: 0,
		CurrentPostLikes:    0,
		Posts:               []types.Post{},
	}

	// Check if the URL path is not root, return not found template
	if r.URL.Path != "/" {
		tmpl, err := template.ParseGlob("ui/templates/notFound.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, r)
		return
	}

	// Get all categories for the topics sidebar
	categories, err := category.GetCategories()
	if err != nil {
		log.Println(err)
	}
	data.Categories = categories

	postID := r.URL.Query().Get("post")
	categorySlug := r.URL.Query().Get("category")

	if postID != "" {
		category, err := category.GetCurrentCategory(categorySlug)

		if err != nil || category.Id == 0 {
			renderNotFoundTemplate(w, r)
			return
		}
		data.CurrentCategory = category
		currentPost, err := post.GetPostById(postID)
		dislikes, likes, err := postRating.GetPostRatings(postID)
		data.CurrentPostDislikes = dislikes
		data.CurrentPostLikes = likes

		if err != nil {
			renderNotFoundTemplate(w, r)
			return
		}

		data.CurrentPost = currentPost
		renderTemplate("ui/templates/post.html", w, data)

	} else if categorySlug != "" {
		category, err := category.GetCurrentCategory(categorySlug)

		if err != nil || category.Id == 0 {
			renderNotFoundTemplate(w, r)
			return
		}

		data.CurrentCategory = category
		posts, err := post.GetCategoryPosts(category)

		if err != nil || len(posts) == 0 {
			log.Println(err)
		}

		data.Posts = posts
		renderTemplate("ui/templates/home.html", w, data)
	} else {
		renderTemplate("ui/templates/home.html", w, data)
	}
}

func renderNotFoundTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("ui/templates/notFound.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, r)
}

func renderTemplate(templatePath string, w http.ResponseWriter, data interface{}) {
	tmpl, err := template.ParseGlob(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, data)
}
