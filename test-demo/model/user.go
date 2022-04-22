package model

type User struct {
	Id           string  `json:"id"`
	SId          string  `json:"sId" xorm:"sId"`
	Name         string  `json:"name"`
	Introduction string  `json:"introduction"`
	Balance      float64 `json:"balance"`
	Icon         string  `json:"icon"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email"`
	// ......其他属性
}
