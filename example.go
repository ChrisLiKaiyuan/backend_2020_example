package main

import (
	"github.com/ChrisLiKaiyuan/backend_2020_example/db"
	"github.com/ChrisLiKaiyuan/backend_2020_example/web"
	_ "gorm.io/driver/mysql"
	log "unknwon.dev/clog/v2"
)

func init() {
	_ = log.NewConsole()
}

func main() {
	//fmt.Println("hello world")

	db.ConnDB()

	web.Run()

}
