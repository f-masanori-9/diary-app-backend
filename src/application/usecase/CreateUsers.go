package usecase

import (
	"diary-app-backend/application/domain/model"
	"diary-app-backend/application/domain/repository"
	"fmt"
)

type CreateUserUseCase struct {
	userRepository repository.IUserRepository
}

func NewCreateUserUseCase(userRepository repository.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}

func (c *CreateUserUseCase) Execute(email, uid string) *model.User {
	user := model.CreateUser(email, uid)
	fmt.Println("CreateUserUseCase.Execute")
	err := c.userRepository.Create(*user)
	if err != nil {
		fmt.Println(err)
	}

	// DBに保存する処理を書く
	return user
}
