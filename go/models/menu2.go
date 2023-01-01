package models

type Menu2 struct {
	Path  string `json:"path"`
	Label string `gorm:"primary_key" json:"label"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Url   string `json:"url"`
}

func (User) Menu2() string {
	return "menu2"
}
