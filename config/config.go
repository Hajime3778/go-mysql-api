package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// DataBaseConfigList foo
type DataBaseConfigList struct {
	Host     string
	User     string
	Password string
	Database string
}

// DataBaseConfig foo
var DataBaseConfig DataBaseConfigList

func init() {
	cfg, err := ini.Load("config.ini")

	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	DataBaseConfig = DataBaseConfigList{
		User:     cfg.Section("DBConnection").Key("user").String(),
		Password: cfg.Section("DBConnection").Key("password").String(),
		Host:     cfg.Section("DBConnection").Key("host").String(),
		Database: cfg.Section("DBConnection").Key("database").String(),
	}
}
