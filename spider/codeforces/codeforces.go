package codeforces

import "github.com/imdario/mergo"

//ScrapeAll 获得所有结果
func ScrapeAll(uid string) (map[string]int, error) {
	res := map[string]int{}
	mainRet, err := GetIntMsg(uid)
	if err != nil {
		return nil, err
	}
	err = mergo.Merge(&res, mainRet)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func ScrapeInt(uid string) (map[string]int, error) {
	mainRet, err := GetIntMsg(uid)
	if err != nil {
		return nil, err
	}
	return mainRet, nil
}
func ScrapeStr(uid string) (map[string]string, error) {
	mainRet, err := GetStrMsg(uid)
	if err != nil {
		return nil, err
	}
	return mainRet, nil
}
