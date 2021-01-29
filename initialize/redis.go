package initialize

import (
	"edushiwei-service/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.COURSE_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.COURSE_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.COURSE_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.COURSE_REDIS = client
	}
}
