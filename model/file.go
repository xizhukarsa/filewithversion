package model

import "gorm.io/gorm"

type (
	File struct {
		Base
		Name      string `gorm:"size:20; comment:'文件名'" json:"name"`
		URL       string `gorm:"size:300; comment:'文件地址'" json:"url"`
		VersionID int64  `gorm:"comment:'版本id'" json:"versionId"`
	}
	FileQuery struct {
		QueryPage
		VersionIDs []int64
	}
)

func (u *File) Create(db *gorm.DB) error {
	return nil
}

func (q FileQuery) List(db *gorm.DB) (*[]File, error) {
	return nil, nil
}
