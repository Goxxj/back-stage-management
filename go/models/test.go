package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Username       string `json:"username"`
	Password       string `json:"password"`
	PrivilegeLevel int    `json:"privilegeLevel"`
}

func (User) TableName2() string {
	return "test"
}
