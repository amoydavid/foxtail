package user

import (
	"foxtail/model"
	"foxtail/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	// 获取当前登录用户
	user := c.MustGet("user").(model.User)

	response.Success(c, http.StatusOK, gin.H{
		"profile": user,
	})
}
