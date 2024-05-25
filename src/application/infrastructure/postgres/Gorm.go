package postgres

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateGormInstanceByConnection(postgresDb *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgresDb,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("gormのインスタンス生成に失敗しました", err)
		panic(err)
	}
	return gormDB
}
