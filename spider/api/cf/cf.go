package cf

import "XCPCer_board/scraper"

func ScrapeAll(uid string) (map[string]string, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, string](
		InitMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
