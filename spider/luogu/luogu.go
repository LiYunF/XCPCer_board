package luogu

import "XCPCer_board/scraper"

func ScrapeAll(uid string) (map[string]int64, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int64](
		GetStrMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
