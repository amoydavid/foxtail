package model

import (
	"database/sql"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PersonalAccessToken struct {
	gorm.Model
	Token      string `gorm:"varchar(120);index"`
	Device     string `gorm:"varchar(120)"`
	UserId     int
	ExpiredAt  sql.NullTime
	LastUsedAt sql.NullTime
	Ability    datatypes.JSON `gorm:"serializer:json"`
	User       User           `gorm:"foreignKey:UserId"`
}
