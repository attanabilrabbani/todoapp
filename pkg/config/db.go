package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func DBConnect() {
	addr := "user={username} password={password} dbname={dbname} sslmode=disable"

	dbConn, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(err)
	}
	db = dbConn

}

func GetDB() *gorm.DB {
	return db
}
