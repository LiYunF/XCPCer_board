package scraper

import (
	"XCPCer_board/model"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scrape 爬
func (s *Scraper[V]) Scrape(url string) (map[string]V, error) {
	ch := make(chan Result[V])
	r := request[V]{
		Url: url,
		Ch:  ch,
	}
	// 发送任务到处理协程
	s.ch <- r
	select {
	case ret := <-ch:
		return ret.getMap(), ret.getError()
	case <-time.After(s.timeout):
		return nil, model.ScrapeTimeoutError
	}
}
