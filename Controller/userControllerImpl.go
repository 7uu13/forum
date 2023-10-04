package Controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/7uu13/forum/Service"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
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
