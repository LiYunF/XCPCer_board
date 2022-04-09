package scraper

import (
	"sync"
)

// @Author: Feng
// @Date: 2022/4/8 22:03

//Result 传输的结果结构
type Result[V any] struct {
	mpMutex  sync.RWMutex
	errMutex sync.RWMutex
	mp       map[string]V
	err      error
}

func NewResultMap[V any]() Result[V] {
	return Result[V]{
		mpMutex:  sync.RWMutex{},
		errMutex: sync.RWMutex{},
		mp:       map[string]V{},
	}
}

//Get 获取结果
func (r *Result[V]) Get(key string) V {
	r.mpMutex.RLock()
	defer r.mpMutex.RUnlock()
	v := r.mp[key]
	return v
}

//Set 设置结果
func (r *Result[V]) Set(key string, value V) {
	r.mpMutex.Lock()
	defer r.mpMutex.Unlock()
	r.mp[key] = value
}

//SetError 设置错误
func (r *Result[V]) SetError(err error) {
	if err == nil {
		return
	}
	r.errMutex.Lock()
	defer r.errMutex.Unlock()
	r.err = err
}

//getError 内部解构方法：获取错误
func (r *Result[V]) getError() error {
	r.errMutex.RLock()
	defer r.errMutex.RUnlock()
	return r.err
}

//getMap 内部解构方法：获取map
func (r *Result[V]) getMap() map[string]V {
	return r.mp
}
