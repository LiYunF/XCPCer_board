package scraper

import (
	"github.com/gocolly/colly"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scraper colly封装
type Scraper[V any] struct {
	c  *colly.Collector
	cb func(*colly.Request) V
	ch chan Request[V]
}

//参数丰富接口
type scraperFunc[V any] func(*Scraper[V])

type Request[V any] struct {
	Url string
	Ch  chan V
}

//NewScraper 构造Scraper
func NewScraper[V any](opts ...scraperFunc[V]) *Scraper[V] {
	s := Scraper[V]{
		c: colly.NewCollector(
			colly.Async(true),
			colly.MaxDepth(1),
		),
		ch: make(chan Request[V]),
	}
	for _, f := range opts {
		if f != nil {
			f(&s)
		}
	}
	s.init()
	return &s
}

//init 初始化
func (s *Scraper[V]) init() {
	ch := make(chan V)
	s.c.OnRequest(func(request *colly.Request) {
		ret := s.cb(request)
		ch <- ret
	})
	go func() {
		for p := range s.ch {
			s.c.Visit(p.Url)
			ret := <-ch
			p.Ch <- ret
		}
	}()
}

//Scrape 爬
func (s *Scraper[V]) Scrape(url string) (V, error) {
	ch := make(chan V)
	r := Request[V]{
		Url: url,
		Ch:  ch,
	}
	s.ch <- r
	select {
	case ret := <-ch:
		return ret, nil
	case <-time.After(5 * time.Second):

	}
	var emptyV V
	return emptyV, nil
}
