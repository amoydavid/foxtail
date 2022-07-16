package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name    string `gorm:"varchar(100);index"`
	AdminId int    `gorm:"int"`
	Admin   User
}
