package mModel

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password"`

	Score uint `json:"score"`
}
