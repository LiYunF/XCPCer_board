package dao

import (
	"XCPCer_board/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

const redisDriver = "redis"

// 初始化连接
func init() {
	// 获取配置
	redisConfig, ok := config.Conf.Storages[redisDriver]
	if !ok {
		panic(fmt.Errorf("lack of redis config"))
	}
	// 初始化
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Open Redis Error:%v", err)
		panic(err)
	}
}
