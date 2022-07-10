package global

import (
	"gblog/config"

	"go.uber.org/zap"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
)
