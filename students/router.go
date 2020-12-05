package students

import (
	. "github.com/ChrisLiKaiyuan/backend_2020_example/toolkit"
	"github.com/gin-gonic/gin"
)

func Router(student *gin.RouterGroup) {
	// add
	student.POST("", Entry(AddStudentInfo))

	// query
	student.GET("", Entry(GetStudentInfo))

	// modify
	student.PUT("", Entry(UpdateStudentInfo))

	// delete
	student.DELETE("", Entry(DeleteStudentInfo))
}
