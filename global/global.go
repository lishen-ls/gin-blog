package global

import (
	"gin-blog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config config.Configuration
	Log    *zap.Logger
	DB     *gorm.DB
}

var App = new(Application)
