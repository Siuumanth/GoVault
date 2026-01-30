package dao

import model "auth/internal/model"

type UserDAO interface {
	// user creation doesnt mean user logging in, just create
	CreateUser(user model.NewUser) error
	GetUserByEmail(email string) (model.DomainUser, error)
}
