package scraper

import (
	"github.com/gocolly/colly"
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

//Processor 处理单元
type Processor struct {
	c  *colly.Collector
	ch chan KV
}

//NewProcessor 新处理单元
func NewProcessor(c *colly.Collector, cb func(collector *colly.Collector, res *Processor)) *Processor {
	p := &Processor{
		ch: make(chan KV, maxKVBufferSize),
		c:  c,
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
