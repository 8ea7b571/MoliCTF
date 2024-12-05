package mModel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Gender   uint
	Phone    string
	Email    string
	Avatar   string
	Birthday time.Time

	Username string
	Password string
	Active   bool
}
