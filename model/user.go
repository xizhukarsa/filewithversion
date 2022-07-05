package model

import "gorm.io/gorm"

type (
	User struct {
		Base
		Name    string `gorm:"size:20; comment:'名字'" json:"name"`
		Token   string `gorm:"size:200; comment:'token'" json:"token"`
		IsAdmin bool   `gorm:"comment:'是否是管理员'" json:"isAdmin"`
	}
	UserQuery struct {
		QueryPage
	}
)

func (u *User) Create(db *gorm.DB) error {
	return nil
}

func (q UserQuery) List(db *gorm.DB) (*[]User, error) {
	return nil, nil
}
