package scraper

import (
	"XCPCer_board/model"
	"github.com/gocolly/colly"
	"time"
)

// @Author: Feng
// @Date: 2022/4/8 18:16

//startListen 启动一个监听者
func (s *Scraper[V]) startListen(collector *colly.Collector, ch chan Result[V]) {
	for p := range s.ch {
		err := collector.Visit(p.Url)
		// 阻塞等待返回结果
		ret := Result[V]{}
		select {
		case ret = <-ch:
			break
		case <-time.After(5 * time.Second):
			err = model.ScrapeTimeoutError
		}
		// 错误判断
		if err != nil {
			ret.Err = err
		}
		p.Ch <- ret
	}
}
