package models

type OutputModel struct {
	StaffID   string `json:"staffID"`
	StaffName string `json:"staffName"`
	Phone     string `json:"phone"`
}

type CreateModel struct {
	StaffID   string `json:"staffID" validate:"required,min=8,max=9"`
	StaffName string `json:"staffName" validate:"required"`
	Phone     string `json:"phone" validate:"required,len=11"`
	//BirthDay  string `json:"birthDay"`
}

type UpdateModel struct {
	StaffName string `json:"staffName"`
	Phone     string `json:"phone" validate:"len=11"`
}
