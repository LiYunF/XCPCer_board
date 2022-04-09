package scraper

import (
	"XCPCer_board/model"
	"fmt"
	"github.com/gocolly/colly"
	"sync"
	"time"
)

// @Author: Feng
// @Date: 2022/4/8 17:38

var s sync.RWMutex

//Scraper colly封装
type Scraper[V any] struct {
	cb      func(collector *colly.Collector, ch chan Result[V])
	ch      chan request[V]
	threads uint32
	timeout time.Duration
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

//WithTimeout 带上并发处理的超时时间
func WithTimeout[V any](timeout time.Duration) scraperFunc[V] {
	if timeout < time.Second {
		timeout = time.Second
	}
	return func(s *Scraper[V]) {
		s.timeout = timeout
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
		timeout: 5 * time.Second,
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
	return &s, nil
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
		go s.newThread(c, ch)
	}
	return nil
}

//newThread 启动一个监听者
func (s *Scraper[V]) newThread(collector *colly.Collector, ch chan Result[V]) {
	for p := range s.ch {
		err := collector.Visit(p.Url)
		// 阻塞等待返回结果
		ret := Result[V]{}
		select {
		case ret = <-ch:
			break
		case <-time.After(s.timeout):
			err = model.ScrapeTimeoutError
		}
		// 错误判断
		if err != nil {
			ret.SetError(err)
		}
		p.Ch <- ret
	}
}
