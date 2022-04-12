package nowcoder

import (
	"XCPCer_board/scraper"
)

// @Author: Feng
// @Date: 2022/4/8 17:09

//ScrapeAll 拉取牛客的所有结果
func ScrapeAll(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		FetchMainPage(uid),
		FetchPractice(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
