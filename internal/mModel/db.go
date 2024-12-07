package mModel

import (
	"errors"
	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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

	err = mdb.initDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: remove this
	mdb.insertTestData()

	return mdb
}

func (mdb *MDB) initDatabase() error {
	return mdb.db.AutoMigrate(
		&Admin{},
		&Challenge{},
		&Flag{},
		&Team{},
		&User{},
	)
}

func (mdb *MDB) insertTestData() {
	admin := Admin{
		Name:     "喻灵",
		Gender:   1,
		Phone:    "13333333333",
		Email:    "admin@qq.com",
		Avatar:   "https://yvling.cn/img/logo.jpeg",
		Birthday: utils.ParseTime("2002-01-01"),
		Username: "yvling",
		Password: "123456",
		Active:   true,
	}

	user := User{
		Name:     "喻灵",
		Gender:   1,
		Phone:    "13333333333",
		Email:    "admin@qq.com",
		Avatar:   "https://yvling.cn/img/logo.jpeg",
		Birthday: utils.ParseTime("2002-01-01"),
		Username: "yvling",
		Password: "123456",
		Active:   true,
		Score:    10000,
		TeamId:   1,
	}

	mdb.CreateAdmin(&admin)
	mdb.CreateUser(&user)
}
