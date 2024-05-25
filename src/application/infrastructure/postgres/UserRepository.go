package postgres

import (
	"diary-app-backend/application/domain/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DBClient *gorm.DB
}

func NewUserRepository(DBClient *gorm.DB) *UserRepository {
	return &UserRepository{DBClient: DBClient}
}

func (u *UserRepository) Create(user model.User) error {
	fmt.Println("UserRepository.Create")
	// result, err := u.DBClient.Exec("INSERT INTO users (email, uid) VALUES ($1, $2)", user.Email, user.UID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	
	return nil
}
