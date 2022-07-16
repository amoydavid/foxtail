package controller

import (
	"errors"
	"foxtail/forms"
	"foxtail/global"
	"foxtail/helper"
	"foxtail/model"
	"foxtail/response"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
)

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		color.Blue(err.Error())
		response.Err(c, http.StatusInternalServerError, -1, err.Error())
		return
	}

	var user = model.User{}

	result := global.DB.Where("email = ?", PasswordLoginForm.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Err(c, http.StatusNotFound, -1, "Error Email or Password")
		return
	}

	if !helper.ValidatePasswords(PasswordLoginForm.Password, user.Password) {
		response.Err(c, http.StatusUnauthorized, -1, "Error Email or Password")
		return
	}

	token, err := user.LoginToken("PC", true)

	if err == nil && len(token) > 0 {
		response.Success(c, http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	response.Err(c, http.StatusInternalServerError, -1, err.Error())
}
