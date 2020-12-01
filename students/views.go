package students

import (
	"fmt"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/db"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/models"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"net/http"
)

func AddStudentInfo(c *gin.Context) {
	// 获取数据
	createModel := new(CreateModel)
	err := c.ShouldBindJSON(createModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "can't bind json",
		})
		return
	}
	fmt.Println("body:", createModel)

	// 数据校验
	if createModel.StaffName == "" || createModel.StaffID == "" || len(createModel.StaffID) != 8 {
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "invalid data",
		})
		return
	}

	// 插入数据库
	student := new(StudentInfoModel)
	student = &StudentInfoModel{
		//Model:     gorm.Model{},
		StaffID:   createModel.StaffID,
		StaffName: createModel.StaffName,
		Phone:     createModel.Phone,
	}
	err = DB.Create(student).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "database wrong",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessReturn{
		Msg:   "success",
		Data:  fmt.Sprintf("%s created", student.StaffID),
		Error: 0,
	})
	return
}

func GetStudentInfo(c *gin.Context) {
	staffID := c.Query("id")
	fmt.Println(staffID)
	if staffID == "" {
		models := make([]*StudentInfoModel, 0, 100)
		DB.Find(&models)

		outputModels := make([]*OutputModel, 0, 100)
		for _, v := range models {
			outputModels = append(outputModels, &OutputModel{
				StaffID:   v.StaffID,
				StaffName: v.StaffName,
				Phone:     v.Phone,
			})
		}

		c.JSON(http.StatusOK, SuccessReturn{
			Msg:   "success",
			Data:  outputModels,
			Error: 0,
		})
	} else {
		Student := new(StudentInfoModel)
		DB.Where("staff_id = ?", staffID).Find(&Student)
		if Student.StaffID == "" {
			c.JSON(http.StatusInternalServerError, MakeErrorReturn{
				Code: 50000,
				Msg:  "id not found or deleted",
			})
		} else {
			var outputModel = OutputModel{
				StaffID:   Student.StaffID,
				StaffName: Student.StaffName,
				Phone:     Student.Phone,
			}
			c.JSON(http.StatusOK, SuccessReturn{
				Msg:   "success",
				Data:  outputModel,
				Error: 0,
			})
		}
	}

	return
}

func UpdateStudentInfo(c *gin.Context) {
	staffID := c.Query("id")
	if staffID == "" {
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "id not provide",
		})
		return
	}
	Student := new(StudentInfoModel)
	_ = DB.Where("staff_id = ?", staffID).Find(&Student)
	if Student.StaffID == "" {
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "id not found or deleted",
		})
		return
	}
	updateModel := new(UpdateModel)
	err := c.ShouldBindJSON(updateModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "can't bind json",
		})
		return
	}
	fmt.Println("body:", updateModel)
	student := new(StudentInfoModel)
	student = &StudentInfoModel{
		StaffName: updateModel.StaffName,
		Phone:     updateModel.Phone,
	}
	student.StaffName = updateModel.StaffName
	student.Phone = updateModel.Phone
	err = DB.Where("staff_id = ?", staffID).Model(&StudentInfoModel{}).Updates(&student).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "database wrong",
		})
	} else {
		c.JSON(http.StatusOK, SuccessReturn{
			Msg:   "success",
			Data:  fmt.Sprintf("%s updated", staffID),
			Error: 0,
		})
	}
	return
}

func DeleteStudentInfo(c *gin.Context) {
	staffID := c.Query("id")
	if staffID == "" {
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "id not provide",
		})
		return
	}
	Student := new(StudentInfoModel)
	_ = DB.Where("staff_id = ?", staffID).Find(&Student)
	if Student.StaffID == "" {
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "id not found or deleted",
		})
		return
	}

	err := DB.Delete(&StudentInfoModel{}, "staff_id = ?", staffID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "database wrong",
		})
	} else {
		c.JSON(http.StatusOK, SuccessReturn{
			Msg:   "success",
			Data:  fmt.Sprintf("%s deleted", staffID),
			Error: 0,
		})
	}
	return
}
