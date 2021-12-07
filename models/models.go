package models

import (
	"gorm.io/gorm"
)

type Employ struct {
	gorm.Model

	Empid   string  `gorm:"primary_key;auto_increment" json:"empid"`
	EmpName string  `json:"empname"`
	DeptID  string  `json :"deptid"`
	Salary  float64 `json :"salary"`
	Email   string  `gorm:"typevarchar(100);unique_index" json:"email"`
	Phone   int64   `json:"phone"`
	Address string  `json:"address"`
}
