package redis

import (
	"XCPCer_board/config"
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var Db *redis.Client

// 初始化连接
func init() {
	red := config.Config.Database.RedisConf
	Db = redis.NewClient(&redis.Options{
		Addr:     red.Host,
		Password: red.Password, // no password set
		DB:       0,            // use default DB
	})
	_, err := Db.Ping(context.TODO()).Result()
	if err != nil {
		log.Errorf("Open Redis Error:%v", err)
	}
}
