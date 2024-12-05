package mModel

import (
	"testing"

	"github.com/8ea7b571/MoliCTF/config"
)

func TestInitDatabase(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mdb := NewMDB()
	err := mdb.InitDatabase()
	if err != nil {
		t.Error(err)
	}
}
