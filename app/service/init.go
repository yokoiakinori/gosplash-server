package service

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"

	"gosplash-server/app/model"
)

var DbEngine *xorm.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	driverName := "mysql"
	DsName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	
	err = errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
