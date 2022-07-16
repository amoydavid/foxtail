package router

import (
	"foxtail/app/http/controller"
	"foxtail/middlewares"
	"foxtail/model"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)

		// 使用身份验证的中间件
		authorized := UserRouter.Group("/", middlewares.Sanctum())

		authorized.GET("list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		authorized.GET("info", func(context *gin.Context) {
			// 获取当前登录用户
			user := context.MustGet("user").(model.User)

			context.JSON(200, gin.H{
				"user": gin.H{
					"id": user.ID,
				},
				"message": "user_info",
			})
		})

	}
}
