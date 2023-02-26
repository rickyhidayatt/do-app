package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/do-app/model"
)

type UserRepository interface {
	GetUserById(id string) ([]model.User, error)
}
type userRepository struct {
	db *sqlx.DB
}

func (r *userRepository) GetUserById(id string) ([]model.User, error) {
	var user []model.User
	err := r.db.Select(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
