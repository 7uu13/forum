package dto

import (
	"github.com/7uu13/forum/model"
)

type UserDTO struct {
    Username string
	FirstName string 
	LastName string
    Email    string
}

// NewUserDTO creates a UserDTO from a User, excluding sensitive fields.
func NewUserDTO(user model.User) UserDTO {
    return UserDTO {
        Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
        Email:    user.Email,
    }
}