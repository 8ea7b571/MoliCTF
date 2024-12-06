package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("D:\\Projects\\Go\\MoliCTF\\config.yaml")
	fmt.Printf("%+v\n", MConfig)
}
