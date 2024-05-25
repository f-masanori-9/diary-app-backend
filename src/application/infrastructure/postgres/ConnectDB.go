package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	*sql.DB
}

func ConnectDB() *sql.DB {
	dataSource := "user=postgres password=password dbname=app sslmode=disable"

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
