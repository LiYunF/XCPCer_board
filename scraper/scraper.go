package scraper

import (
	"XCPCer_board/model"
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scraper colly封装
type Scraper[V any] struct {
	c  *colly.Collector
	cb func(collector *colly.Collector, ch chan V)
	ch chan request[V]
}

//参数丰富接口
type scraperFunc[V any] func(*Scraper[V])

//WithCallback 带上处理回调函数
func WithCallback[V any](cb func(collector *colly.Collector, ch chan V)) scraperFunc[V] {
	return func(s *Scraper[V]) {
		s.cb = cb
	}
}

func defaultCallback[V any]() func(collector *colly.Collector, ch chan V) {
	return func(c *colly.Collector, ch chan V) {
		c.OnRequest(func(req *colly.Request) {
			fmt.Println(req.URL)
		})
		var v V
		ch <- v
	}
}

type request[V any] struct {
	Url string
	Ch  chan result[V]
}

type result[V any] struct {
	Value V
	err   error
}

//NewScraper 构造Scraper
func NewScraper[V any](opts ...scraperFunc[V]) (*Scraper[V], error) {
	s := Scraper[V]{
		c: colly.NewCollector(
			colly.Async(true),
			colly.MaxDepth(1),
		),
		ch: make(chan request[V]),
		cb: defaultCallback[V](),
	}
	for _, f := range opts {
		if f != nil {
			f(&s)
		}
	}
	err := s.init()
	if err != nil {
		return nil, err
	}
	return &s, err
}

//init 初始化
func (s *Scraper[V]) init() error {
	// 初始化各种On
	ch := make(chan V)
	s.cb(s.c, ch)

	// 开始监听
	go func() {
		for p := range s.ch {
			err := s.c.Visit(p.Url)
			ret := <-ch
			p.Ch <- result[V]{Value: ret, err: err}
		}
	}()
	return nil
}

//Scrape 爬
func (s *Scraper[V]) Scrape(url string) (V, error) {
	var emptyV V
	ch := make(chan result[V])
	r := request[V]{
		Url: url,
		Ch:  ch,
	}
	s.ch <- r
	select {
	case ret := <-ch:
		return ret.Value, ret.err
	case <-time.After(5 * time.Second):
		return emptyV, model.ScrapeTimeoutError
	}
	return emptyV, nil
}
