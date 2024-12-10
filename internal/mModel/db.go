package mModel

import (
	"errors"
	"github.com/8ea7b571/MoliCTF/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type MDB struct {
	Type string

	db *gorm.DB
}

func NewMDB() *MDB {
	var err error

	mdb := new(MDB)
	mdb.Type = config.MConfig.MDatabase.Type

	switch mdb.Type {
	case "sqlite":
		mdb.db, err = gorm.Open(sqlite.Open(config.MConfig.MApp.Root+"/moli.db"), &gorm.Config{})
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
	for i := 0; i < 16; i++ {
		admin := Admin{
			Name:     "喻灵",
			Gender:   1,
			Phone:    "13333333333",
			Email:    "admin@qq.com",
			Avatar:   "https://yvling.cn/img/logo.jpeg",
			Username: "yvling",
			Password: "123456",
			Active:   true,
		}

		team := Team{
			Name:        "Test team",
			Description: "just for test",
			Avatar:      "https://yvling.cn/img/logo.jpeg",
			Password:    "123456",
			MemberNum:   111,
			Score:       20000,
		}

		user := User{
			Name:         "yvling",
			Gender:       1,
			Phone:        "13333333333",
			Email:        "admin@qq.com",
			Avatar:       "https://yvling.cn/img/logo.jpeg",
			Introduction: "Fuck you!",
			Username:     "yvling",
			Password:     "123456",
			Active:       true,
			Score:        10000,
			TeamId:       1,
		}

		mdb.CreateAdmin(&admin)
		mdb.CreateUser(&user)
		mdb.CreateTeam(&team)
	}
}
