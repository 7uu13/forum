package Controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/7uu13/forum/Service"
	"github.com/7uu13/forum/Model"
)

type UserControllerImpl struct {
	userService Service.UserService
}

func NewUserController(userService Service.UserService) UserController {
	return &UserControllerImpl{userService}
}

func (c *UserControllerImpl) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid User Id", http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetUserByID(userId)
	if err != nil {
		http.Error(w, "User Not Found!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}

func (c *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    userID, err := c.userService.CreateUser(user)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    response := map[string]int64{"id": userID}
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}