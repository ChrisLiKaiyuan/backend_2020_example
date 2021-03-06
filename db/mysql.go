package db

import (
	"fmt"
	"github.com/ChrisLiKaiyuan/backend_2020_example/conf"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func initDatabase() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level [Silent Error Warn Info]
			Colorful:      true,          // 禁用彩色打印
		},
	)

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Get().MySQL.User,
		conf.Get().MySQL.Password,
		conf.Get().MySQL.Addr,
		conf.Get().MySQL.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("init database success")
	return db
}

func ConnDB() {
	DB = initDatabase()
	err := DB.AutoMigrate(
		&StudentInfoModel{},
	)
	if err != nil {
		panic(err)
	}
}
