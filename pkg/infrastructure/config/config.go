package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// DataBaseConfigList foo
type DataBaseConfigList struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Config config
type Config struct {
	Server struct {
		Port    string
		Timeout int
	}
	DataBase struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

// NewConfig create config
func NewConfig() *Config {

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	c := new(Config)

	// conf読み取り
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// UnmarshalしてConfigにマッピング
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return c
}
