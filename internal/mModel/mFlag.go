package mModel

import "gorm.io/gorm"

type Flag struct {
	gorm.Model

	UserID      uint
	ChallengeID uint
	FlagValue   string
}
