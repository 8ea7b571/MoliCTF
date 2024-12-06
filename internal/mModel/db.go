package mModel

import (
	"errors"
	"gorm.io/driver/sqlite"
	"log"

	"github.com/8ea7b571/MoliCTF/config"
	"gorm.io/gorm"
)

type MDB struct {
	Type string
	Path string

	db *gorm.DB
}

func NewMDB() *MDB {
	var err error

	mdb := new(MDB)
	mdb.Type = config.MConfig.MDatabase.Type
	mdb.Path = config.MConfig.MDatabase.Path

	switch mdb.Type {
	case "sqlite":
		mdb.db, err = gorm.Open(sqlite.Open(mdb.Path), &gorm.Config{})
		break
	default:
		log.Fatal(errors.New("unknown database type"))
	}

	if err != nil {
		log.Fatal(err)
	}

	return mdb
}

func (mdb *MDB) InitDatabase() error {
	return mdb.db.AutoMigrate(
		&Admin{},
	)
}
