package model

import "gorm.io/gorm"

type (
	Space struct {
		Base
		Name string `gorm:"size:20; comment:'空间名字'" json:"name"`
		Spec string `gorm:"size:200; comment:'描述信息'" json:"spec"`
	}
	SpaceQuery struct {
		QueryPage
	}
)

func (s *Space) Create(db *gorm.DB) error {
	return nil
}

func (q SpaceQuery) List(db *gorm.DB) (*[]Space, error) {
	return nil, nil
}
