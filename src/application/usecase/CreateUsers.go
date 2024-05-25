package usecase

import (
	"fmt"
	"my-learning-box-backend-v2/application/domain/model"
	"my-learning-box-backend-v2/application/domain/repository"
)

type CreateUserUseCase struct {
	userRepository repository.IUserRepository
}

func NewCreateUserUseCase(userRepository repository.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}



func (c *CreateUserUseCase) Execute(email, uid string) *model.User {
	user := model.CreateUser(email, uid)
	err := c.userRepository.Create(*user)
	fmt.Println(err)
	// DBに保存する処理を書く
	return user
}
