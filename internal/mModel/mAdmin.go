package mModel

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
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

func (mdb *MDB) CreateAdmin(admin *Admin) (int64, error) {
	result := mdb.db.Create(admin)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) DeleteAdmin(admin *Admin) (int64, error) {
	result := mdb.db.Delete(admin)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) UpdateAdmin(admin *Admin) (int64, error) {
	result := mdb.db.Save(admin)
	return result.RowsAffected, result.Error
}

func (mdb *MDB) GetAdminWithId(id int64) (*Admin, error) {
	admin := &Admin{}
	result := mdb.db.First(admin, id)
	return admin, result.Error
}
