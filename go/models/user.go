package models

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PrivilegeLevel int    `gorm:"default:2;" json:"privilegeLevel"`
}

func (User) TableName() string {
	return "user"
}
