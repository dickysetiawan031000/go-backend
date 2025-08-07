package mapper

import (
	"github.com/dickysetiawan031000/go-backend/dto/auth"
	"github.com/dickysetiawan031000/go-backend/model"
)

func ToUserResponse(user model.User) auth.UserResponse {
	return auth.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
