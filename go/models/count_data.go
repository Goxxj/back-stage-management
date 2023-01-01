package models

type CountData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

func (User) CountData() string {
	return "count_data"
}
