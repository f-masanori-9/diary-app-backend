package postgres

import (
	"database/sql"
	"my-learning-box-backend-v2/application/domain/model"
)

type UserRepository struct {
	DBClient *sql.DB
}

func NewUserRepository(DBClient *sql.DB) *UserRepository {
	return &UserRepository{DBClient: DBClient}
}

func (u *UserRepository) Create(user model.User) error {
	return nil
}
