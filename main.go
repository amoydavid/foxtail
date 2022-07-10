package main

import (
	"fmt"
	"gblog/global"
	"gblog/initialize"
	"gblog/router"

	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	r := router.SetupRouter()

	//r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))

	zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
}
