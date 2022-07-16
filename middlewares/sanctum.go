package middlewares

import (
	"errors"
	"foxtail/global"
	"foxtail/helper"
	"foxtail/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 类似于 Laravel 中的身份验证，token经Sha256后存于personal_access_token表
// header里放 Authorization: Bearer 1234567
func Sanctum() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getBearerTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		var personalAccessToken model.PersonalAccessToken

		result := global.DB.Preload("User").Where("token = ? and expired_at > ?", helper.Sha256(token), time.Now()).First(&personalAccessToken)

		if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		personalAccessToken.LastUsedAt.Time = time.Now()
		personalAccessToken.LastUsedAt.Valid = true
		personalAccessToken.ExpiredAt.Time = time.Now().Add(time.Hour * 24 * 14)
		personalAccessToken.ExpiredAt.Valid = true
		global.DB.Save(&personalAccessToken)

		c.Set("user", personalAccessToken.User)

		c.Next()
	}
}

// getBearerTokenFromHeader returns the token from the header
func getBearerTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("not found authorization header")
	}

	token := strings.Split(header, " ")
	if len(token) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	if token[0] != "Bearer" {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return token[1], nil
}
