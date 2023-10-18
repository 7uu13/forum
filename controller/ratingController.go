package controller

import (
	"net/http"
	"strconv"

	"github.com/7uu13/forum/types"
)

type RatingController struct{}

func (_ *RatingController) RatingController(w http.ResponseWriter, r *http.Request) {
	var postRating types.PostRating
	var replyRating types.ReplyRating

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
		rating_id_string := r.URL.Query().Get("rating_id")

		// Define a common function to process ratings.
		processRating := func(id int, rating string, handleFunc func(int, int, string)) {
			if id == 0 {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}

			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			ratingValue := r.FormValue("rating")
			if user_id == 0 || ratingValue == "" {
				http.Error(w, "Missing parameters", http.StatusBadRequest)
				return
			}

			handleFunc(id, user_id, ratingValue)
		}

		if post_id_string != "" {

			postID, err := strconv.Atoi(post_id_string)
			if err != nil {
				http.Error(w, "Invalid post_id", http.StatusBadRequest)
				return
			}
			processRating(postID, "rating", postRating.HandlePostRating)
			http.Redirect(w, r, "/?post="+post_id_string, http.StatusSeeOther)

		} else if rating_id_string != "" {
			referer := r.Header.Get("Referer")
			ratingID, err := strconv.Atoi(rating_id_string)
			if err != nil {
				http.Error(w, "Invalid rating_id", http.StatusBadRequest)
				return
			}
			processRating(ratingID, "rating", replyRating.HandleReplyRating)
			http.Redirect(w, r, referer, http.StatusSeeOther)
		}
	}
}
