package students

import (
	"fmt"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/db"
	. "github.com/ChrisLiKaiyuan/backend_2020_example/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	log "unknwon.dev/clog/v2"
)

func AddStudentInfo(c *gin.Context) (int, int, interface{}) {
	// 获取数据
	createModel := new(CreateModel)
	err := c.ShouldBindJSON(createModel)
	log.Trace("body: %+v", createModel)
	if err != nil {
		log.Warn("can't bind json")
		return http.StatusBadRequest, 40000, "can't bind json"
	}

	// 数据校验
	validate := validator.New()
	err = validate.Struct(createModel)
	if err != nil {
		log.Warn("invalid data")
		return http.StatusBadRequest, 40000, "invalid data"
	}

	// 插入数据库
	student := &StudentInfoModel{
		//Model:     gorm.Model{},
		StaffID:   createModel.StaffID,
		StaffName: createModel.StaffName,
		Phone:     createModel.Phone,
	}
	tx := DB.Begin()
	err = tx.Create(student).Error
	if err != nil {
		tx.Rollback()
		log.Warn("database wrong")
		return http.StatusInternalServerError, 50000, "database wrong"
	}
	tx.Commit()
	log.Trace("%s created", student.StaffID)
	return http.StatusOK, 20000, fmt.Sprintf("%s created", student.StaffID)
}

func GetStudentInfo(c *gin.Context) (int, int, interface{}) {
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
		return http.StatusOK, 20000, outputModels
	} else {
		student := StudentInfoModel{
			StaffID: staffID,
		}
		err := DB.Where(&student).Take(&student).Error
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			log.Trace("id %s not found or deleted", staffID)
			return http.StatusInternalServerError, 50000, fmt.Sprintf("id %s not found or deleted", staffID)
		} else {
			var outputModel = OutputModel{
				StaffID:   student.StaffID,
				StaffName: student.StaffName,
				Phone:     student.Phone,
			}
			log.Trace("Queried data of id %s", staffID)
			return http.StatusOK, 20000, outputModel
		}
	}
}

func UpdateStudentInfo(c *gin.Context) (int, int, interface{}) {
	staffID := c.Query("id")
	if staffID == "" {
		log.Warn("id not provide")
		return http.StatusBadRequest, 40000, "id not provide"
	}
	student := StudentInfoModel{
		StaffID: staffID,
	}
	err := DB.Where(&student).Take(&student).Error
	if err == gorm.ErrRecordNotFound {
		log.Warn("id %s not found or deleted", staffID)
		return http.StatusInternalServerError, 50000, fmt.Sprintf("id %s not found or deleted", staffID)
	}
	updateModel := new(UpdateModel)
	err = c.ShouldBindJSON(updateModel)
	log.Trace("id: %s body: %+v", staffID, updateModel)
	if err != nil {
		log.Warn("can't bind json")
		return http.StatusBadRequest, 40000, "can't bind json'"
	}
	// check
	validate := validator.New()
	err = validate.Struct(updateModel)
	if err != nil {
		log.Warn("invalid data")
		return http.StatusBadRequest, 40000, "invalid data"
	}
	student = StudentInfoModel{
		StaffName: updateModel.StaffName,
		Phone:     updateModel.Phone,
	}
	tx := DB.Begin()
	err = tx.Where(student).Model(&StudentInfoModel{}).Updates(&student).Error
	if err != nil {
		tx.Rollback()
		log.Warn("database wrong")
		return http.StatusInternalServerError, 50000, "database wrong"
	} else {
		tx.Commit()
		log.Trace("%s updated", staffID)
		return http.StatusOK, 20000, fmt.Sprintf("%s updated", staffID)
	}
}

func DeleteStudentInfo(c *gin.Context) (int, int, interface{}) {
	staffID := c.Query("id")
	if staffID == "" {
		log.Warn("id not provide")
		return http.StatusBadRequest, 40000, "id not provide"
	}
	student := StudentInfoModel{
		StaffID: staffID,
	}
	err := DB.Where(&student).Take(&student).Error
	if err == gorm.ErrRecordNotFound {
		log.Warn("id %s not found or deleted", staffID)
		return http.StatusInternalServerError, 50000, fmt.Sprintf("id %s not found or deleted", staffID)
	}
	tx := DB.Begin()
	err = tx.Delete(&StudentInfoModel{}, &student).Error
	if err != nil {
		tx.Rollback()
		log.Warn("database wrong")
		return http.StatusInternalServerError, 50000, "database wrong"
	} else {
		tx.Commit()
		log.Trace("%s deleted", staffID)
		return http.StatusOK, 20000, fmt.Sprintf("%s deleted", staffID)
	}
}
