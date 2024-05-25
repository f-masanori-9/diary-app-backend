package repository

import "my-learning-box-backend-v2/application/domain/model"

type IBaseRepository[T model.User] interface {
	Create(arg T) error

	
}
