package scraper

import (
	"XCPCer_board/model"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scrape 爬
func (s *Scraper[V]) Scrape(url string) ([]KV[V], error) {
	select {
	case p := <-s.ch:
		kvs, err := p.Collect(url)
		s.ch <- p
		return kvs, err
	case <-time.After(s.timeout):
		return nil, model.ScrapeTimeoutError
	}
}
