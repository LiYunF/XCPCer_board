package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"sync"
)

// @Author: Feng
// @Date: 2022/4/8 17:38

var s sync.RWMutex

//Scraper colly封装
type Scraper[V any] struct {
	cb      func(collector *colly.Collector, ch chan Result[V])
	ch      chan request[V]
	threads uint32
}

//request 传输的请求结构
type request[V any] struct {
	Url string
	Ch  chan Result[V]
}

//参数丰富接口
type scraperFunc[V any] func(*Scraper[V])

//WithCallback 带上处理回调函数
func WithCallback[V any](cb func(collector *colly.Collector, ch chan Result[V])) scraperFunc[V] {
	return func(s *Scraper[V]) {
		s.cb = cb
	}
}

//WithThreads 带上并发处理的协程数
func WithThreads[V any](threads uint32) scraperFunc[V] {
	return func(s *Scraper[V]) {
		s.threads = threads
	}
}

//defaultCallback 默认的处理回调函数
func defaultCallback[V any]() func(collector *colly.Collector, ch chan Result[V]) {
	return func(c *colly.Collector, ch chan Result[V]) {
		c.OnRequest(func(req *colly.Request) {
			fmt.Println(req.URL)
		})
		ch <- Result[V]{}
	}
}

//NewScraper 构造Scraper
func NewScraper[V any](opts ...scraperFunc[V]) (*Scraper[V], error) {
	// 默认参数
	s := Scraper[V]{
		threads: 1,
		ch:      make(chan request[V]),
		cb:      defaultCallback[V](),
	}
	// 应用外来参数
	for _, f := range opts {
		if f != nil {
			f(&s)
		}
	}
	// 初始化
	err := s.init()
	if err != nil {
		return nil, err
	}
	return &s, err
}

//init 初始化
func (s *Scraper[V]) init() error {
	// 初始化各种On
	for i := uint32(0); i < s.threads; i++ {
		ch := make(chan Result[V])
		c := colly.NewCollector(
			colly.Async(true),
			colly.MaxDepth(1),
		)
		s.cb(c, ch)
		// 开始监听
		go s.startListen(c, ch)
	}
	return nil
}
