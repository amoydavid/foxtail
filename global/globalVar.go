package global

import (
	"foxtail/config"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
	DB       *gorm.DB
)
