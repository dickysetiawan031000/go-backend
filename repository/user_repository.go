package repository

import "github.com/dickysetiawan031000/go-backend/model"

type UserRepository interface {
	Register(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepoMemory struct {
	users []model.User
}

func NewUserRepository() UserRepository {
	return &userRepoMemory{
		users: []model.User{},
	}
}

func (r *userRepoMemory) Register(user *model.User) error {
	user.ID = uint(int64(len(r.users) + 1))
	r.users = append(r.users, *user)
	return nil
}

func (r *userRepoMemory) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, nil // not found, no error
}
