package lib

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"SatohAyaka/leaving-match-backend/model"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func InitDB() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("DB connect error: %v", err)
		}
		db.AutoMigrate(&model.BusTime{}, &model.Vote{}, &model.Result{})
		DB = db
	})
	return DB
}
