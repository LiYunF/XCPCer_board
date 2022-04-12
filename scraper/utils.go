package scraper

import "github.com/imdario/mergo"

// @Author: Feng
// @Date: 2022/4/12 19:32

type equal interface{ ~string | ~int32 | ~int64 }

//MergeAllResults 合并所有结果
func MergeAllResults[K equal, V any](results ...Results[V]) (map[K]V, error) {
	res, err := make(map[K]V), error(nil)
	for _, i := range results {
		if i.GetError() != nil {
			return nil, err
		}
		err = mergo.Merge(&res, i.GetMap())
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
