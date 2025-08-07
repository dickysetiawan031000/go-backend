package usecase

import (
	"errors"

	"github.com/dickysetiawan031000/go-backend/dto/auth"
	"github.com/dickysetiawan031000/go-backend/model"
	"github.com/dickysetiawan031000/go-backend/utils"
)

type AuthUseCase interface {
	Register(user model.User) (model.User, error)
	Login(email, password string) (model.User, error)
	GetProfile(id uint) (model.User, error)
	UpdateProfile(id uint, input auth.UpdateUserRequest) (model.User, error)
}

type authUseCase struct {
	users []model.User
}

func NewAuthUseCase() AuthUseCase {
	return &authUseCase{
		users: []model.User{},
	}
}

// Register akan membuat user baru dan menyimpan ke slice in-memory
func (a *authUseCase) Register(user model.User) (model.User, error) {
	for _, u := range a.users {
		if u.Email == user.Email {
			return model.User{}, errors.New("email already exists")
		}
	}

	// Hash password sebelum simpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = hashedPassword

	user.ID = uint(len(a.users) + 1)
	a.users = append(a.users, user)
	return user, nil
}

// Login memverifikasi kombinasi email dan password
func (a *authUseCase) Login(email, password string) (model.User, error) {
	for _, u := range a.users {
		if u.Email == email && utils.CheckPasswordHash(password, u.Password) {
			return u, nil
		}
	}
	return model.User{}, errors.New("invalid email or password")
}

func (a *authUseCase) GetProfile(userID uint) (model.User, error) {
	for _, user := range a.users {
		if user.ID == userID {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (a *authUseCase) UpdateProfile(id uint, input auth.UpdateUserRequest) (model.User, error) {
	for i, u := range a.users {
		if u.ID == id {
			// Cek duplikat email
			for _, other := range a.users {
				if other.Email == input.Email && other.ID != id {
					return model.User{}, errors.New("email already taken")
				}
			}

			a.users[i].Name = input.Name
			a.users[i].Email = input.Email

			// Kalau mau ganti password, harus masukkan password lama
			if input.NewPassword != "" {
				if input.OldPassword == "" {
					return model.User{}, errors.New("old password required to change password")
				}

				if !utils.CheckPasswordHash(input.OldPassword, u.Password) {
					return model.User{}, errors.New("old password is incorrect")
				}

				hashed, err := utils.HashPassword(input.NewPassword)
				if err != nil {
					return model.User{}, err
				}
				a.users[i].Password = hashed
			}

			return a.users[i], nil
		}
	}
	return model.User{}, errors.New("user not found")
}
