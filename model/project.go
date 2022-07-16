package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name    string `gorm:"varchar(100);index"`
	AdminId int    `gorm:"int"`
	Admin   User
	Users   []*User `gorm:"many2many:user_projects;"`
}
