package Redis

import (
	"XCPCer_board/config"
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var Rdb *redis.Client

// 初始化连接
func init() {
	red := config.Config.Database.RedisConf
	Rdb = redis.NewClient(&redis.Options{
		Addr:     red.Host,
		Password: red.Password, // no password set
		DB:       0,            // use default DB
	})
	_, err := Rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Errorf("Open Redis Error:%v", err)
	}
}
