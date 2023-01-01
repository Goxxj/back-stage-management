package models

type UserDetails struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Sex   string `json:"sex"`
	Birth string `json:"birth"`
	Addr  string `json:"addr"`
}

func (User) UserDetails() string {
	return "user_details"
}
