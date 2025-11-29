package database

import model "auth/internal/model"

type UserDao interface {
	// user creation doesnt mean user logging in, just create
	CreateUser(user model.NewUser) error
	GetUserByEmail(email string) (model.DomainUser, error)
}
