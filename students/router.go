package students

import (
	"github.com/gin-gonic/gin"
)

func Router(student *gin.RouterGroup) {
	// add
	student.POST("", AddStudentInfo)

	// query
	student.GET("", GetStudentInfo)

	// modify
	student.PUT("", UpdateStudentInfo)

	// delete
	student.DELETE("", DeleteStudentInfo)
}
