package models

type Menu struct {
	Path     string `json:"path"`
	Label    string `gorm:"primary_key" json:"label"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Url      string `json:"url"`
	Children BList  `json:"children"`
}

func (User) Menu() string {
	return "menu"
}
