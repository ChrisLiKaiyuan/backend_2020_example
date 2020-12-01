package models

import "gorm.io/gorm"

type StudentInfoModel struct {
	gorm.Model
	StaffID   string `gorm:"uniqueIndex;size:50"`
	StaffName string
	Phone     string
	//BirthDay  time.Time
}
