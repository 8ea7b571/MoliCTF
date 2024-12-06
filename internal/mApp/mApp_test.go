package mApp

import (
	"testing"

	"github.com/8ea7b571/MoliCTF/config"
)

func TestMApp(t *testing.T) {
	config.LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")

	mapp := NewMApp()
	err := mapp.Run()
	if err != nil {
		t.Error(err)
	}
}
