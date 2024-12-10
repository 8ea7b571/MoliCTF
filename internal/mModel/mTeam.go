package mModel

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password"`

	MemberNum uint `json:"member_num"`
	Score     uint `json:"score"`
}

func (mdb *MDB) CreateTeam(team *Team) (int64, error) {
	result := mdb.db.Create(team)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) GetTeams(offset, limit int) ([]*Team, error) {
	var teams []*Team
	result := mdb.db.Limit(limit).Offset(offset).Find(&teams)
	return teams, result.Error
}

func (mdb *MDB) GetTeamWithId(id uint) (*Team, error) {
	team := &Team{}
	result := mdb.db.First(team, id)
	return team, result.Error
}

func (mdb *MDB) GetTeamCount() (int, error) {
	var count int64
	result := mdb.db.Model(&Team{}).Count(&count)
	return int(count), result.Error
}
