package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Base struct {
		ID       int64     `gorm:"index; autoIncrement; primaryKey; comment:'id'" json:"id"`
		CreateAt time.Time `gorm:"autoCreateTime; comment:'创建时间'" json:"createAt"`
	}

	QueryPage struct {
		Page int
		Size int
	}
	QueryOrder struct {
	}
)

var (
	modelList = []interface{}{
		&User{},
		&Space{},
		&Version{},
		&File{},
	}
)

func AutoMigration(db *gorm.DB) error {
	return db.AutoMigrate(modelList...)
}
