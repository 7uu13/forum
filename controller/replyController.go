package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/7uu13/forum/types"
)

type ReplyController struct{}

func (_ *ReplyController) ReplyController(w http.ResponseWriter, r *http.Request) {
	var postReply types.PostReply

	switch r.Method {
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

		content := r.FormValue("content")
		fmt.Println(content)
		if user_id == 0 || post_id == 0 || content == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}

		postReply.CreatePostReply(post_id, user_id, content)

		http.Redirect(w, r, "/?post="+post_id_string, http.StatusSeeOther)
	}
}
