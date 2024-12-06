package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type mConfig struct {
	MApp      mAppConfig      `yaml:"m_app"`
	MDatabase mDatabaseConfig `yaml:"m_database"`
}

type mAppConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Expire   int    `yaml:"expire"`
	Template string `yaml:"template"`
}

type mDatabaseConfig struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

var MConfig mConfig

func LoadConfig(configPath string) {
	cFile, err := os.OpenFile(configPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer cFile.Close()

	cBytes, err := io.ReadAll(cFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(cBytes, &MConfig)
	if err != nil {
		log.Fatal(err)
	}
}
