package students

import (
	"fmt"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/db"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	_ "gorm.io/driver/mysql"
	"net/http"
	log "unknwon.dev/clog/v2"
)

func AddStudentInfo(c *gin.Context) {
	// 获取数据
	createModel := new(CreateModel)
	err := c.ShouldBindJSON(createModel)
	log.Trace("body: %+v", createModel)
	if err != nil {
		log.Warn("can't bind json")
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "can't bind json",
		})
		return
	}

	// 数据校验
	validate := validator.New()
	err = validate.Struct(createModel)
	if err != nil {
		log.Warn("invalid data")
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
		log.Warn("database wrong")
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "database wrong",
		})
		return
	}

	log.Trace("%s created", student.StaffID)
	c.JSON(http.StatusOK, SuccessReturn{
		Msg:   "success",
		Data:  fmt.Sprintf("%s created", student.StaffID),
		Error: 0,
	})
	return
}

func GetStudentInfo(c *gin.Context) {
	staffID := c.Query("id")
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

		log.Trace("Queried all data")
		c.JSON(http.StatusOK, SuccessReturn{
			Msg:   "success",
			Data:  outputModels,
			Error: 0,
		})
	} else {
		Student := new(StudentInfoModel)
		DB.Where("staff_id = ?", staffID).Find(&Student)
		if Student.StaffID == "" {
			log.Trace("id %s not found or deleted", staffID)
			c.JSON(http.StatusInternalServerError, MakeErrorReturn{
				Code: 50000,
				Msg:  fmt.Sprintf("id %s not found or deleted", staffID),
			})
		} else {
			var outputModel = OutputModel{
				StaffID:   Student.StaffID,
				StaffName: Student.StaffName,
				Phone:     Student.Phone,
			}
			log.Trace("Queried data of id %s", staffID)
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
		log.Warn("id not provide")
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "id not provide",
		})
		return
	}
	Student := new(StudentInfoModel)
	_ = DB.Where("staff_id = ?", staffID).Find(&Student)
	if Student.StaffID == "" {
		log.Warn("id %s not found or deleted", staffID)
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  fmt.Sprintf("id %s not found or deleted", staffID),
		})
		return
	}
	updateModel := new(UpdateModel)
	err := c.ShouldBindJSON(updateModel)
	log.Trace("id: %s body: %+v", staffID, updateModel)
	if err != nil {
		log.Warn("can't bind json")
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "can't bind json",
		})
		return
	}
	// check
	validate := validator.New()
	err = validate.Struct(updateModel)
	if err != nil {
		log.Warn("invalid data")
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "invalid data",
		})
		return
	}
	student := new(StudentInfoModel)
	student = &StudentInfoModel{
		StaffName: updateModel.StaffName,
		Phone:     updateModel.Phone,
	}
	student.StaffName = updateModel.StaffName
	student.Phone = updateModel.Phone
	err = DB.Where("staff_id = ?", staffID).Model(&StudentInfoModel{}).Updates(&student).Error
	if err != nil {
		log.Warn("database wrong")
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "database wrong",
		})
	} else {
		log.Trace("%s updated", staffID)
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
		log.Warn("id not provide")
		c.JSON(http.StatusBadRequest, MakeErrorReturn{
			Code: 40000,
			Msg:  "id not provide",
		})
		return
	}
	Student := new(StudentInfoModel)
	_ = DB.Where("staff_id = ?", staffID).Find(&Student)
	if Student.StaffID == "" {
		log.Warn("id %s not found or deleted", staffID)
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  fmt.Sprintf("id %s not found or deleted", staffID),
		})
		return
	}

	err := DB.Delete(&StudentInfoModel{}, "staff_id = ?", staffID).Error
	if err != nil {
		log.Warn("database wrong")
		c.JSON(http.StatusInternalServerError, MakeErrorReturn{
			Code: 50000,
			Msg:  "database wrong",
		})
	} else {
		log.Trace("%s deleted", staffID)
		c.JSON(http.StatusOK, SuccessReturn{
			Msg:   "success",
			Data:  fmt.Sprintf("%s deleted", staffID),
			Error: 0,
		})
	}
	return
}
