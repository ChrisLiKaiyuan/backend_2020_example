package web

import (
	"fmt"
	"github.com/ChrisLiKaiyuan/backend_2020_example/conf"
	"github.com/ChrisLiKaiyuan/backend_2020_example/students"
	"github.com/gin-gonic/gin"
)

func Routers(routers *gin.Engine) {
	student := routers.Group("/student")
	students.Router(student)
}

func Run() {
	r := gin.Default()
	Routers(r)

	err := r.Run(fmt.Sprintf(":%d", conf.Get().Port))
	if err != nil {
		panic(err)
	}
}
