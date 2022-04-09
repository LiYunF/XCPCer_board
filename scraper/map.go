package scraper

import "sync/atomic"

// @Author: Feng
// @Date: 2022/4/8 22:03

//resultMap Scrape结果map
type resultMap[V any] struct {
	mutexFlag int32
	mp        map[string]V
}

func (r *resultMap[V]) Get(key string) V {
	for {
		atomic.CompareAndSwapInt64()
	}
}
