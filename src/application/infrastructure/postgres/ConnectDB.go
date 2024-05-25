package postgres

import (
	"database/sql"
	"diary-app-backend/config"
	"fmt"
	"log"

	"time"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	*sql.DB
}

func ConnectDB(config *config.Config) *sql.DB {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		*config.APIPort,
		config.DB.User,
		config.DB.Password,
		config.DB.DBName)

	var db *sql.DB
	var err error
	maxRetries := 5
	retryInterval := 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", dataSource)
		if err == nil {
			break
		}

		log.Printf("データベースの接続に失敗しました。リトライします... (エラー: %v)", err)

		time.Sleep(retryInterval)
	}
	db.Ping()
	if err != nil {
		log.Print("リトライ後もデータベースの接続に失敗しました。")
		return nil
	} else {
		fmt.Println("データベースに接続しました。")
	}

	return db
}
