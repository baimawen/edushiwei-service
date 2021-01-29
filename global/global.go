package global

import (
	"edushiwei-service/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	COURSE_DB     *gorm.DB
	COURSE_REDIS  *redis.Client
	COURSE_CONFIG config.Server
	COURSE_VP     *viper.Viper
	COURSE_LOG    *zap.Logger
)
