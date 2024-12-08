package mModel

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name   string `json:"name"`
	Gender uint   `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`

	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`

	Score  uint `json:"score"`
	TeamId uint `json:"team_id"`
}

func (mdb *MDB) CreateUser(user *User) (int64, error) {
	result := mdb.db.Create(user)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) DeleteUser(user *User) (int64, error) {
	result := mdb.db.Delete(user)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) UpdateUser(user *User) (int64, error) {
	result := mdb.db.Save(user)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) GetUserWithId(id uint) (*User, error) {
	user := &User{}
	result := mdb.db.First(user, id)
	return user, result.Error
}

func (mdb *MDB) GetUserWithUsername(username string) (*User, error) {
	user := &User{}
	result := mdb.db.Where("username = ?", username).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return user, nil
}
