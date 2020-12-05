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
		//c.JSON(http.StatusBadRequest, ErrorReturn{
		//	Code: 40000,
		//	Msg:  "can't bind json",
		//})
		return http.StatusBadRequest, 40000, "can't bind json"
	}

	// 数据校验
	validate := validator.New()
	err = validate.Struct(createModel)
	if err != nil {
		log.Warn("invalid data")
		//c.JSON(http.StatusBadRequest, ErrorReturn{
		//	Code: 40000,
		//	Msg:  "invalid data",
		//})
		return http.StatusBadRequest, 40000, "invalid data"
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
		//c.JSON(http.StatusInternalServerError, ErrorReturn{
		//	Code: 50000,
		//	Msg:  "database wrong",
		//})
		return http.StatusInternalServerError, 50000, "database wrong"
	}

	log.Trace("%s created", student.StaffID)
	//c.JSON(http.StatusOK, SuccessReturn{
	//	Msg:   "success",
	//	Data:  fmt.Sprintf("%s created", student.StaffID),
	//	Error: 0,
	//})
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
		//c.JSON(http.StatusOK, SuccessReturn{
		//	Msg:   "success",
		//	Data:  outputModels,
		//	Error: 0,
		//})
		return http.StatusOK, 20000, outputModels
	} else {
		Student := new(StudentInfoModel)
		err := DB.Where("staff_id = ?", staffID).Find(&Student).Error
		if err == gorm.ErrRecordNotFound {
			log.Trace("id %s not found or deleted", staffID)
			//c.JSON(http.StatusInternalServerError, ErrorReturn{
			//	Code: 50000,
			//	Msg:  fmt.Sprintf("id %s not found or deleted", staffID),
			//})
			return http.StatusInternalServerError, 50000, fmt.Sprintf("id %s not found or deleted", staffID)
		} else {
			var outputModel = OutputModel{
				StaffID:   Student.StaffID,
				StaffName: Student.StaffName,
				Phone:     Student.Phone,
			}
			log.Trace("Queried data of id %s", staffID)
			//c.JSON(http.StatusOK, SuccessReturn{
			//	Msg:   "success",
			//	Data:  outputModel,
			//	Error: 0,
			//})
			return http.StatusOK, 20000, outputModel
		}
	}
}

func UpdateStudentInfo(c *gin.Context) (int, int, interface{}) {
	staffID := c.Query("id")
	if staffID == "" {
		log.Warn("id not provide")
		//c.JSON(http.StatusBadRequest, ErrorReturn{
		//	Code: 40000,
		//	Msg:  "id not provide",
		//})
		return http.StatusBadRequest, 40000, "id not provide"
	}
	Student := new(StudentInfoModel)
	err := DB.Where("staff_id = ?", staffID).Find(&Student).Error
	if err == gorm.ErrRecordNotFound {
		log.Warn("id %s not found or deleted", staffID)
		//c.JSON(http.StatusInternalServerError, ErrorReturn{
		//	Code: 50000,
		//	Msg:  fmt.Sprintf("id %s not found or deleted", staffID),
		//})
		return http.StatusInternalServerError, 50000, fmt.Sprintf("id %s not found or deleted", staffID)
	}
	updateModel := new(UpdateModel)
	err = c.ShouldBindJSON(updateModel)
	log.Trace("id: %s body: %+v", staffID, updateModel)
	if err != nil {
		log.Warn("can't bind json")
		//c.JSON(http.StatusBadRequest, ErrorReturn{
		//	Code: 40000,
		//	Msg:  "can't bind json",
		//})
		return http.StatusBadRequest, 40000, "can't bind json'"
	}
	// check
	validate := validator.New()
	err = validate.Struct(updateModel)
	if err != nil {
		log.Warn("invalid data")
		//c.JSON(http.StatusBadRequest, ErrorReturn{
		//	Code: 40000,
		//	Msg:  "invalid data",
		//})
		return http.StatusBadRequest, 40000, "invalid data"
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
		//c.JSON(http.StatusInternalServerError, ErrorReturn{
		//	Code: 50000,
		//	Msg:  "database wrong",
		//})
		return http.StatusInternalServerError, 50000, "database wrong"
	} else {
		log.Trace("%s updated", staffID)
		//c.JSON(http.StatusOK, SuccessReturn{
		//	Msg:   "success",
		//	Data:  fmt.Sprintf("%s updated", staffID),
		//	Error: 0,
		//})
		return http.StatusOK, 20000, fmt.Sprintf("%s updated", staffID)
	}
}

func DeleteStudentInfo(c *gin.Context) (int, int, interface{}) {
	staffID := c.Query("id")
	if staffID == "" {
		log.Warn("id not provide")
		//c.JSON(http.StatusBadRequest, ErrorReturn{
		//	Code: 40000,
		//	Msg:  "id not provide",
		//})
		return http.StatusBadRequest, 40000, "id not provide"
	}
	Student := new(StudentInfoModel)
	err := DB.Where("staff_id = ?", staffID).Find(&Student).Error
	if err == gorm.ErrRecordNotFound {
		log.Warn("id %s not found or deleted", staffID)
		//c.JSON(http.StatusInternalServerError, ErrorReturn{
		//	Code: 50000,
		//	Msg:  fmt.Sprintf("id %s not found or deleted", staffID),
		//})
		return http.StatusInternalServerError, 50000, fmt.Sprintf("id %s not found or deleted", staffID)
	}

	err = DB.Delete(&StudentInfoModel{}, "staff_id = ?", staffID).Error
	if err != nil {
		log.Warn("database wrong")
		//c.JSON(http.StatusInternalServerError, ErrorReturn{
		//	Code: 50000,
		//	Msg:  "database wrong",
		//})
		return http.StatusInternalServerError, 50000, "database wrong"
	} else {
		log.Trace("%s deleted", staffID)
		//c.JSON(http.StatusOK, SuccessReturn{
		//	Msg:   "success",
		//	Data:  fmt.Sprintf("%s deleted", staffID),
		//	Error: 0,
		//})
		return http.StatusOK, 20000, fmt.Sprintf("%s deleted", staffID)
	}
}
