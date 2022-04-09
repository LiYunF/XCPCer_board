package scraper

import (
	"sync"
)

// @Author: Feng
// @Date: 2022/4/8 22:03

//Result 传输的结果结构
type Result[V any] struct {
	m   sync.RWMutex
	mp  map[string]V
	Err error
}

func NewResultMap[V any]() Result[V] {
	return Result[V]{
		m:  sync.RWMutex{},
		mp: map[string]V{},
	}
}

//Get 获取结果
func (r *Result[V]) Get(key string) V {
	r.m.RLock()
	defer r.m.RUnlock()
	v := r.mp[key]
	return v
}

//Set 设置结果
func (r *Result[V]) Set(key string, value V) {
	r.m.Lock()
	defer r.m.Unlock()
	r.mp[key] = value
}
