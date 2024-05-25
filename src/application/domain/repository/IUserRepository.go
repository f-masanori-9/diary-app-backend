package repository

import "diary-app-backend/application/domain/model"

type IUserRepository interface {
	Create(user model.User) error
}
