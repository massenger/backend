package storage

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ConnectDB ...
func ConnectDB(username, password, server, dbname string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + server + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic("database couldn't connect. see error above")
	}
	initializeTable(db)
	return db
}

func initializeTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
	log.Println("sql table `User` succeeded")
}
