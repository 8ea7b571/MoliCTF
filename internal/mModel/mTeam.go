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

func (mdb *MDB) CreateTeam(team *Team) (int64, error) {
	result := mdb.db.Create(team)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) GetTeamWithId(id uint) (*Team, error) {
	team := &Team{}
	result := mdb.db.First(team, id)
	return team, result.Error
}
