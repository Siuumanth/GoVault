package database

import (
	model "auth/internal/model"
	"database/sql"
)

// now write implementations for PGSQL
type PGUserDAO struct {
	db *sql.DB
}

// init pg user DAO struct
func NewPostgresUserDAO(db *sql.DB) *PGUserDAO {
	return &PGUserDAO{db: db}
}

func (p *PGUserDAO) CreateUser(user model.NewUser) error {
	// we have CreateUserQuery
	// db.exec when we dont expect rows back
	_, err := p.db.Exec(CreateUserQuery, user.Email, user.Username, user.PasswordHash)
	return err
}

func (p *PGUserDAO) GetUserByEmail(email string) (model.DomainUser, error) {
	// we have GetUserByEmailQuery
	rows := p.db.QueryRow(email)

	var user model.DomainUser
	err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.DomainUser{}, nil // no user found
		}
		return model.DomainUser{}, err // some error
	}
	return user, nil
}
