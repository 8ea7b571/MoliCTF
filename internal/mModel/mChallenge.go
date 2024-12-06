package mModel

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description" gorm:"type:text"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	ConnInfo    string `json:"conn_info"`
	InitScore   uint   `json:"init_score"`
	MiniScore   uint   `json:"mini_score"`
	Visible     bool   `json:"visible"`
}
