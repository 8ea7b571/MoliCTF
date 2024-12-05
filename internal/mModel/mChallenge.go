package mModel

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model

	Name        string
	Description string `gorm:"type:text"`
	Category    string
	Image       string
	ConnInfo    string
	InitScore   uint
	MiniScore   uint
	Visible     bool
}
