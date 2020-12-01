package models

type OutputModel struct {
	StaffID   string `json:"staffID"`
	StaffName string `json:"staffName"`
	Phone     string `json:"phone"`
}

type CreateModel struct {
	StaffID   string `json:"staffID"`
	StaffName string `json:"staffName"`
	Phone     string `json:"phone"`
	//BirthDay  string `json:"birthDay"`
}

type UpdateModel struct {
	StaffName string `json:"staffName"`
	Phone     string `json:"phone"`
}
