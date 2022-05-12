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
type KV struct {
	Key string
	Val interface{}
}

//GetPersistHandler 获取持久化处理函数
func (k *KV) GetPersistHandler(cb *PersistHandler) Persist {
	return func() error {
		return cb.Do(k.Key, k.Val)
	}
}

//Processor 处理单元
type Processor struct {
	c     *colly.Collector
	ch    chan KV
	errCh chan error
}

//NewProcessor 新处理单元
func NewProcessor(c *colly.Collector, cb func(collector *colly.Collector, res *Processor)) *Processor {
	p := &Processor{
		errCh: make(chan error, maxKVBufferSize),
		ch:    make(chan KV, maxKVBufferSize),
		c:     c,
	}
	cb(c, p)
	return p
}

//Set 设置值
func (r *Processor) Set(key string, value interface{}) {
	r.ch <- KV{
		Key: key,
		Val: value,
	}
}

//SetError 设置错误
func (r *Processor) SetError(err error) {
	r.errCh <- err
}

//collect 收集所有返回结果
func (r *Processor) collect(url string) (re []KV, err error) {
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
