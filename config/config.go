package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Server struct {
		Port int
	}
	DB struct {
		FilePath string
	}
}

func LoadConfig(cnfFile string) (*Config, error) {
	return nil, nil
}

func (c *Config) ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(c.DB.FilePath), &gorm.Config{})
}
