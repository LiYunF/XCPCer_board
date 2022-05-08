package luogu

import (
	"XCPCer_board/dao"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func getUserIDRedisKey(uid string, key any) string {
	return fmt.Sprintf("luogu_id_%v_%v", uid, key)
}

//mSet 事务处理set
func mSet(mapKey map[string]int, ctx context.Context) error {

	pipe := dao.RedisClient.TxPipeline()
	for key, val := range mapKey {
		pipe.Set(ctx, key, val, 0)
	}
	_, err := pipe.Exec(ctx)
	return err
}

//SetUserMsgToRedis 将用户信息放入redis
func SetUserMsgToRedis(uid string, ctx context.Context) error {

	//get user msg
	res, err := ScrapeUser(uid)
	if err != nil {
		log.Errorf("get luogu uid=%v message err:%v", uid, err)
	}

	//creat key map
	mapKey := make(map[string]int)
	for key, val := range res {
		mapKey[getUserIDRedisKey(uid, key)] = val
	}

	//set data to redis
	err = mSet(mapKey, ctx)
	if err != nil {
		log.Errorf("set redis data for uid=%v failed, err:%v\n", uid, err)
		return err
	}
	return nil
}

//GetUserMsgFromRedis 从redis中获取用户信息，return -1,err 证明出现错误（可能没找到数据）
func GetUserMsgFromRedis(uid string, keyWord string, ctx context.Context) (int, error) {
	val, err := dao.RedisClient.Get(ctx, getUserIDRedisKey(uid, keyWord)).Result()
	if err != nil {
		return -1, err
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Errorf("luogu redis strToInt get err:%v\tand the return is %v:", num, err)
		return -1, err
	}
	return num, nil
}

func GetUserAllMsgFromRedis(uid string, ctx context.Context) (map[string]int, error) {

}
