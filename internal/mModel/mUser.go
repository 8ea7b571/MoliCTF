package mModel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string    `json:"name"`
	Gender   uint      `json:"gender"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`

	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}
