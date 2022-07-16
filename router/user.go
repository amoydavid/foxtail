package router

import (
	"foxtail/app/http/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})

		UserRouter.POST("login", controller.PasswordLogin)
	}
}
