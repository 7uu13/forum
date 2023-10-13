package Controller

import "net/http"

type PostController interface {
	CreatePostHandler(w http.ResponseWriter, r *http.Request)
}
