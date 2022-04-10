package scraper

import "sync"

// @Author: Feng
// @Date: 2022/4/8 22:03

//Results 结果集
type Results[V any] struct {
	m   sync.Mutex
	mp  map[string]V
	err error
}

func NewResults[V any]() *Results[V] {
	return &Results[V]{
		m:  sync.Mutex{},
		mp: make(map[string]V),
	}
}

//init 使用后重新初始化结果集
func (r *Results[V]) init() {
	r.mp = make(map[string]V)
	r.err = nil
}

//GetMp 获取结果集中的map
func (r *Results[V]) GetMp() map[string]V {
	if r == nil {
		return nil
	}
	return r.mp
}

//GetErr 获取结果集中的error
func (r *Results[V]) GetErr() error {
	if r == nil {
		return nil
	}
	return r.err
}

//Set 设置值
func (r *Results[V]) Set(key string, value V) {
	r.m.Lock()
	defer r.m.Unlock()
	r.mp[key] = value
}

//SetError 设置错误
func (r *Results[V]) SetError(err error) {
	r.m.Lock()
	defer r.m.Unlock()
	r.err = err
}
