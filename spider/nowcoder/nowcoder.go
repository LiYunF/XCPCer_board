package nowcoder

import (
	"XCPCer_board/scraper"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/8 17:09

var (
	// 空持久化
	emptyHandler = scraper.NewPersistHandler[int](emptyPersistHandler[int])
	// 爬取函数
	fetchers = []func(uid string) ([]scraper.KV[int], error){
		FetchMainPage,
		FetchPractice,
	}
	// 匹配持久化处理函数
	persistHandlerMap = map[string]*scraper.PersistHandler[int]{
		mainRatingKey:              emptyHandler,
		mainRatingRatingKey:        emptyHandler,
		mainAttendContestAmountKey: emptyHandler,
		practicePassAmountKey:      emptyHandler,
	}
)

//scrape 拉取牛客的所有结果
func scrape(uid string) (res []scraper.KV[int]) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		kvs, err := f(uid)
		if err != nil {
			log.Errorf("GetPersistHandler Fetcher Error %v", err)
			continue
		}
		res = append(res, kvs...)
	}

	return res
}

//emptyPersistHandler 空持久化函数
func emptyPersistHandler[V any](key string, val V) error {
	//dao.RedisClient.Set()
	//dao.DBClient.ExecContext()
	log.Infof("Nowcoder Key %v Val %v", key, val)
	return nil
}

//matchPersistHandlers 匹配持久化函数
func matchPersistHandlers(kvs []scraper.KV[int]) []scraper.Persist {
	var res []scraper.Persist
	for ind, kv := range kvs {
		h, ok := persistHandlerMap[kvs[ind].Key]
		if ok {
			log.Infof("Get Handler Key %v val %v", kv.Key, kv.Val)
			res = append(res, kvs[ind].GetPersistHandler(h))
		}
	}
	return res
}

//Flush 刷新某用户牛客id信息
func Flush(uid string) {
	// 拉出所有kv对
	kvs := scrape(uid)
	// 为所有key对匹配持久化函数
	persists := matchPersistHandlers(kvs)
	// 向持久化处理协程注册持久化处理函数
	scraper.RegisterPersist(persists...)
}
