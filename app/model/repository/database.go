package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func init() {
	loadEnv()

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_CONTAINER_ADDRESS")
	dbName := os.Getenv("DB_DATABASE")

	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		dbUser, dbPassword, dbAddress, dbName
	)

	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込みに失敗しました。")
	}
}