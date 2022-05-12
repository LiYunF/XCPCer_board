package scraper

import (
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
	cb      func(collector *colly.Collector, res *Processor[V])
	ch      chan *Processor[V]
	threads uint32
	timeout time.Duration
}

//参数丰富接口
type scraperFunc[V any] func(*Scraper[V])

//WithCallback 带上处理回调函数
func WithCallback[V any](cb func(collector *colly.Collector, res *Processor[V])) scraperFunc[V] {
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
func defaultCallback[V any]() func(*colly.Collector, *Processor[V]) {
	return func(c *colly.Collector, res *Processor[V]) {
		c.OnRequest(func(r *colly.Request) {
			fmt.Println(r.URL)
			var v V
			res.Set("Default Callback 1", v)
		})
	}
}

//NewScraper 构造Scraper
func NewScraper[V any](opts ...scraperFunc[V]) *Scraper[V] {
	// 默认参数
	s := Scraper[V]{
		timeout: 5 * time.Second,
		threads: 1,
		cb:      defaultCallback[V](),
	}
	// 应用外来参数
	for _, f := range opts {
		if f != nil {
			f(&s)
		}
	}
	s.ch = make(chan *Processor[V], s.threads)
	// 初始化
	s.init()
	return &s
}

//init 初始化
func (s *Scraper[V]) init() {
	wg := sync.WaitGroup{}
	// 初始化各种On
	for i := uint32(0); i < s.threads; i++ {
		// 初始化并开始监听
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := colly.NewCollector(
				colly.Async(false),
				colly.MaxDepth(1),
				colly.AllowURLRevisit(),
			)
			p := NewProcessor[V](c, s.cb)
			s.ch <- p
		}()
		// 启动持久化处理携程
		go newPersistProcessor()
	}
	wg.Wait()
}
