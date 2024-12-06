package mModel

import (
	"fmt"
	"github.com/8ea7b571/MoliCTF/config"
	"github.com/8ea7b571/MoliCTF/utils"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mdb := NewMDB()
	err := mdb.initDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestMDB_CreateAdmin(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mdb := NewMDB()
	admin := Admin{
		Name:     "喻灵",
		Gender:   1,
		Phone:    "13333333333",
		Email:    "admin@qq.com",
		Avatar:   "https://yvling.cn/img/logo.jpeg",
		Birthday: utils.ParseTime("2002-01-22"),
		Username: "yvling",
		Password: "123456",
		Active:   true,
	}

	affectedRows, err := mdb.CreateAdmin(&admin)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(affectedRows)
}

// the 'Delete' method does not actually delete the data,
// but records the time when the data was deleted.
func TestMDB_DeleteAdmin(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mdb := NewMDB()
	admin, err := mdb.GetAdminWithId(1)
	if err != nil {
		t.Error(err)
	}

	affectedRows, err := mdb.DeleteAdmin(admin)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(affectedRows)
}

func TestMDB_UpdateAdmin(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mdb := NewMDB()
	admin, err := mdb.GetAdminWithId(1)
	if err != nil {
		t.Error(err)
	}

	admin.Birthday = utils.ParseTime("2002-02-25")
	affectedRows, err := mdb.UpdateAdmin(admin)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(affectedRows)
}

func TestMDB_GetAdminWithId(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mdb := NewMDB()
	admin, err := mdb.GetAdminWithId(1)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", admin)
}
