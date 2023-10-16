package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/7uu13/forum/types"
)

type RatingController struct{}

func (_ *RatingController) RatingController(w http.ResponseWriter, r *http.Request) {
	var postRating types.PostRating
	/*
		GET - Gets ratings for a post and returns them as JSON
		POST - Creates a rating for a post
		PUT - Updates a rating for a post
		DELETE - Deletes a rating for a post
	*/
	switch r.Method {
	case "GET":

	case "POST":
		user_id := 12345
		post_id_string := r.URL.Query().Get("post_id")

		post_id, err := strconv.Atoi(post_id_string)
		if err != nil {
			post_id = 0
		}

		err = r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		rating := r.FormValue("rating")
		fmt.Println(rating)
		if user_id == 0 || post_id == 0 || rating == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}

		postRating.HandlePostRating(post_id, user_id, rating)

		http.Redirect(w, r, "/?post="+post_id_string, http.StatusSeeOther)
	}

}
