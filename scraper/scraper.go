package scraper

import (
	"XCPCer_board/model"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scrape 爬
func (s *Scraper) Scrape(url string) ([]KV, error) {
	select {
	case p := <-s.ch:
		kvs, err := p.collect(url)
		s.ch <- p
		return kvs, err
	case <-time.After(s.timeout):
		return nil, model.ScrapeTimeoutError
	}
}
