package models

type TableData struct {
	Name     string `json:"name"`
	TodayBuy int    `json:"todayBuy"`
	MonthBuy int    `json:"monthBuy"`
	TotalBuy int    `json:"totalBuy"`
}

func (User) TableData() string {
	return "table_data"
}
