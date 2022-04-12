package scraper

import (
	"XCPCer_board/model"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scrape 爬
func (s *Scraper[V]) Scrape(url string) Results[V] {
	ch := make(chan Results[V])
	r := request[V]{
		Url: url,
		Ch:  ch,
	}
	// 发送任务到处理协程
	s.ch <- r
	select {
	case ret := <-ch:
		return ret
	case <-time.After(s.timeout):
		return NewResultsWithError[V](model.ScrapeTimeoutError)
	}
}
