package scraper

import (
	"XCPCer_board/model"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// @Author: Feng
// @Date: 2022/4/8 17:38

var s sync.RWMutex

//Scraper colly封装
type Scraper[V any] struct {
	cb      func(collector *colly.Collector, res *Results[V])
	ch      chan request[V]
	threads uint32
	timeout time.Duration
}

//request 传输的请求结构
type request[V any] struct {
	Url string
	Ch  chan Results[V]
}

//参数丰富接口
type scraperFunc[V any] func(*Scraper[V])

//WithCallback 带上处理回调函数
func WithCallback[V any](cb func(collector *colly.Collector, res *Results[V])) scraperFunc[V] {
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
func defaultCallback[V any]() func(*colly.Collector, *Results[V]) {
	return func(c *colly.Collector, res *Results[V]) {
		c.OnRequest(func(r *colly.Request) {
			fmt.Println(r.URL)
			var v V
			res.Set("Default Callback 1", v)
		})
		c.OnScraped(func(r *colly.Response) {
			fmt.Println(string(r.Body))
			var v V
			res.Set("Default Callback 2", v)
		})
	}
}

//NewScraper 构造Scraper
func NewScraper[V any](opts ...scraperFunc[V]) *Scraper[V] {
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
	s.init()
	return &s
}

//init 初始化
func (s *Scraper[V]) init() {
	// 初始化各种On
	for i := uint32(0); i < s.threads; i++ {
		res := NewResults[V]()
		c := colly.NewCollector(
			colly.Async(false),
			colly.MaxDepth(1),
			colly.AllowURLRevisit(),
		)
		// 初始化并开始监听
		go s.newThread(c, res)
	}
}

//newThread 启动一个监听者
func (s *Scraper[V]) newThread(collector *colly.Collector, res *Results[V]) {
	// 初始化on函数
	s.cb(collector, res)
	finCh := make(chan struct{})
	// 开始阻塞监听
	for p := range s.ch {
		var err error
		go func() {
			err = collector.Visit(p.Url)
			if err != nil {
				log.Errorf("Scraper Visit Error url:%v %v", p.Url, err)
				res.SetError(err)
			}
			finCh <- struct{}{}
		}()
		// 阻塞等待返回结果
		select {
		case <-finCh:
			break
		case <-time.After(s.timeout):
			res.SetError(model.ScrapeTimeoutError)
		}
		// 进行结果返回
		p.Ch <- NewResultsWithMapAndError(res.GetMap(), res.GetError())
		// 重新初始化结果集
		res.init()
	}
}
