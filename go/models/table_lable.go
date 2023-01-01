package models

type TableLable struct {
	Name     string `json:"name"`
	TodayBuy string `json:"todayBuy"`
	MonthBuy string `json:"monthBuy"`
	TotalBuy string `json:"totalBuy"`
}

func (User) TableLable() string {
	return "table_lable"
}
