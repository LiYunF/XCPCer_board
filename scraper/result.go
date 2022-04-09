package scraper

// @Author: Feng
// @Date: 2022/4/8 22:03

//result 传输的结果结构
type result[V any] struct {
	Key   string
	Value V
	Err   error
}

//NewResult 正常返回
func NewResult[V any](key string, value V) result[V] {
	return result[V]{Key: key, Value: value}
}

//NewResultWithErr 错误返回
func NewResultWithErr[V any](err error) result[V] {
	return result[V]{Err: err}
}

//Results 结果集
type Results[V any] struct {
	mp  map[string]V
	err error
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
