package router

import (
	"foxtail/controller/user"
	"foxtail/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", user.PasswordLogin)

		// 使用身份验证的中间件
		authorized := UserRouter.Group("/", middlewares.Sanctum())
		authorized.GET("info", user.Info)

	}
}
