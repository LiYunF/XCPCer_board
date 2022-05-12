package scraper

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/8 22:03

const (
	maxKVBufferSize = 128
)

//KV 处理单元返回的键值对
type KV[V any] struct {
	Key string
	Val V
}

//GetPersistHandler 获取持久化处理函数
func (k *KV[V]) GetPersistHandler(cb *PersistHandler[V]) Persist {
	return func() error {
		return cb.Do(k.Key, k.Val)
	}
}

//Processor 处理单元
type Processor[V any] struct {
	c     *colly.Collector
	ch    chan KV[V]
	errCh chan error
}

//NewProcessor 新处理单元
func NewProcessor[V any](c *colly.Collector, cb func(collector *colly.Collector, res *Processor[V])) *Processor[V] {
	p := &Processor[V]{
		errCh: make(chan error, maxKVBufferSize),
		ch:    make(chan KV[V], maxKVBufferSize),
		c:     c,
	}
	cb(c, p)
	return p
}

//Set 设置值
func (r *Processor[V]) Set(key string, value V) {
	r.ch <- KV[V]{
		Key: key,
		Val: value,
	}
}

//SetError 设置错误
func (r *Processor[V]) SetError(err error) {
	r.errCh <- err
}

//Collect 收集所有返回结果
func (r *Processor[V]) Collect(url string) (re []KV[V], err error) {
	// 等待执行结束
	err = r.c.Visit(url)
	if err != nil {
		log.Errorf("Scraper Error %v", err)
		return nil, err
	}
	fin := false
	for {
		select {
		case kv := <-r.ch:
			re = append(re, kv)
		case err1 := <-r.errCh:
			err = err1
		default:
			fin = true
		}
		if fin {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return re, nil
}
