package model

import "gorm.io/gorm"

type (
	Version struct {
		Base
		SpaceID int64  `gorm:"comment:'空间id'" json:"spaceId"`
		Name    string `gorm:"size:20; comment:'版本名称'" json:"name"`
		Spec    string `gorm:"size:200; comment:'版本描述'" json:"spec"`
	}
	VersionQuery struct {
		QueryPage
		SpaceIDs []int64
	}
)

func (v *Version) Create(db *gorm.DB) error {
	return nil
}

func (q VersionQuery) List(db *gorm.DB) (*[]Version, error) {
	return nil, nil
}
