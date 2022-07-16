package main

import (
	"errors"
	"fmt"
	"foxtail/global"
	"foxtail/helper"
	"foxtail/initialize"
	"foxtail/model"
	"foxtail/router"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()

	r := router.SetupRouter()

	var superAdmin model.User
	result := global.DB.First(&superAdmin, 1)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.DB.Create(&model.User{Username: "admin", Email: "admin@admin.com", Password: helper.HashPassword("123456")})
	}

	//r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	zap.L().Info(fmt.Sprintf("running at PORT: %d", global.Settings.Port))
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))

	zap.L().Info("this is hello func", zap.String("error", "启动错误!"))

	//defer global.DB.
}
