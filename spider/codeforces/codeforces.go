package codeforces

import (
	"XCPCer_board/scraper"
)

//ScrapeAll 获得所有结果
func ScrapeAll(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		GetIntMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//ScrapeInt 获取所有整数结果
func ScrapeInt(uid string) (map[string]int, error) {
	res, err := scraper.MergeAllResults[string, int](
		GetIntMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//ScrapeStr 获取所有字符串结果
func ScrapeStr(uid string) (map[string]string, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, string](
		GetStrMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
