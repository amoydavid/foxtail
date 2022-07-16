package model

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"foxtail/global"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string     `gorm:"varchar(100);index;unique;not null;"`
	Email    string     `gorm:"varchar(100);index;unique;not null;"`
	Password string     `gorm:"varchar(120);not null;"`
	Projects []*Project `gorm:"many2many:user_projects;"`
}

func (u *User) LoginToken(device string, dropOther bool) (string, error) {
	if u.ID > 0 {

		var sum, id string

		if dropOther {
			global.DB.Where("user_id = ?", u.ID).Delete(&PersonalAccessToken{})
		}

		for {
			id, _ = gonanoid.New()
			sum = fmt.Sprintf("%x", sha256.Sum256([]byte(id)))

			var existsToken PersonalAccessToken

			result := global.DB.Where("token = ?", sum).First(&existsToken)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				break
			}
		}

		token := PersonalAccessToken{
			Device: device,
			Token:  sum,
			UserId: int(u.ID),
			ExpiredAt: sql.NullTime{
				Time:  time.Now().Add(time.Hour * 24 * 14),
				Valid: true,
			},
			Ability: datatypes.JSON(`["*"]`),
		}
		result := global.DB.Create(&token)
		if result.Error != nil {
			return "", result.Error
		}

		return id, nil
	}

	return "", errors.New("not found record" + fmt.Sprintf("%d", u.ID))
}
