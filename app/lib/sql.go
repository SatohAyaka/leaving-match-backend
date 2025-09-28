package lib

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

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
		loc, _ := time.LoadLocation("Asia/Tokyo")

		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NowFunc: func() time.Time {
				return time.Now().In(loc)
			},
		})
		if err != nil {
			log.Fatalf("DB connect error: %v", err)
		}
		db.AutoMigrate(&model.User{}, &model.Recommended{}, &model.BusTime{}, &model.Vote{}, &model.Result{})
		DB = db
	})
	return DB
}
