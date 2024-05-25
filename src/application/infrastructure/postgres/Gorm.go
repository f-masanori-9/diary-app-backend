package postgres

import (
	"diary-app-backend/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateGormInstance(conf *config.Config) *gorm.DB {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.DBName,
	)

	gormDB, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		fmt.Println("gormのインスタンス生成に失敗しました")
		panic(err)
	}

	return gormDB
}
