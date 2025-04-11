package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

// Singleton pattern implementation
var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func NewMySQLConnection() *gorm.DB {
	dbOnce.Do(func() {
		dsn := "root:roooooot@tcp(127.0.0.1:3306)/BaseDeDatos?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}
		dbInstance = db
	})
	return dbInstance
}
