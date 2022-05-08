package luogu

import (
	"XCPCer_board/dao"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func getUserIDRedisKey(uid string, key any) string {
	return fmt.Sprintf("luogu_id_%v_%v", uid, key)
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
	err = dao.redisClient.MSet(ctx, mapKey).Err()
	if err != nil {
		log.Errorf("set redis data for uid=%v failed, err:%v\n", uid, err)
		return err
	}
	return nil
}

func GetUserMsgFromRedis() error {

	return nil
}
