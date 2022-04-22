package model

type User struct {
	Id   string `json:"id"`
	SId  string `json:"sId" xorm:"sId"`
	Name string `json:"name"`
}
