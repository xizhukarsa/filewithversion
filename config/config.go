package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

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
	f, err := os.Open(cnfFile)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if nil != err {
		return nil, err
	}

	var cnf Config
	if err := json.Unmarshal(buf, &cnf); nil != err {
		return nil, err
	} else {
		return &cnf, nil
	}
}

func (c *Config) ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(c.DB.FilePath), &gorm.Config{})
}
