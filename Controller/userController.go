package Controller

import (
	"net/http"
)

type UserController interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}
