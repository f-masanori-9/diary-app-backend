package repository

import "my-learning-box-backend-v2/application/domain/model"

type IUserRepository interface {
	Create(user model.User) error
}
