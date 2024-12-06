package mModel

import "gorm.io/gorm"

type Flag struct {
	gorm.Model

	UserID      uint   `json:"user_id"`
	ChallengeID uint   `json:"challenge_id"`
	FlagValue   string `json:"flag_value"`
}
