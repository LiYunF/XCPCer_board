package scraper

import (
	"XCPCer_board/dao"
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/5/12 13:36

var (
	// 持久化任务分配管道
	flushCh = make(chan interface{})
)

type redisRequest struct {
	kvs []KV
}

type dbRequest struct {
	do func(cli *sql.DB) error
}

//newFlushProcessor 新持久化处理器
func newFlushProcessor() {
	for {
		i := <-flushCh
		switch v := i.(type) {
		case *redisRequest:
			internalFlushRedis(v)
		case *dbRequest:
			internalFlushDB(v)
		}
	}
}

//internalFlushRedis 内部刷新redis数据
func internalFlushRedis(req *redisRequest) {
	var args []interface{}
	for _, kv := range req.kvs {
		args = append(args, kv.Key, kv.Val)
	}
	// 底层库实现了自动重试
	err := dao.RedisClient.MSet(context.Background(), args...).Err()
	if err != nil {
		log.Errorf("internal flush redis error %v", err)
	}
}

//internalFlushDB 内部刷新db内数据
func internalFlushDB(req *dbRequest) {
	err := req.do(dao.DBClient)
	if err != nil {
		log.Errorf("internal flush db error %v", err)
	}
}

//FlushRedis 刷新Redis
func FlushRedis(kvs []KV) {
	flushCh <- redisRequest{
		kvs: kvs,
	}
}

//FlushDB 刷新DB
func FlushDB(callback func(cli *sql.DB) error) error {
	flushCh <- dbRequest{
		do: callback,
	}
}
