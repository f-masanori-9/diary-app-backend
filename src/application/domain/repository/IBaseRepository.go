package repository

import "diary-app-backend/application/domain/model"

type IBaseRepository[T model.User] interface {
	Create(arg T) error
}
