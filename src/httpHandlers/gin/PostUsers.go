package gin

import (
	"diary-app-backend/application/domain/model"
	"diary-app-backend/application/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PostUsers struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewPostUsers(createUserUseCase *usecase.CreateUserUseCase) *PostUsers {
	return &PostUsers{createUserUseCase: createUserUseCase}
}

func (p *PostUsers) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := model.CreateUser("email", "uid")
		err := p.createUserUseCase.Execute(user.Email, user.UID)
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	}
}
