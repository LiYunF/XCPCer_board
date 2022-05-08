package atcoder

import (
	"XCPCer_board/scraper"
)

//ScrapeAll 拉取所有结果
func ScrapeAll(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		FetchMainPage(uid),
	)
	if err != nil {
		return nil, err
	}

	allCid, _ := ScrapeCid(uid)

	//fmt.Println(allCid)

	for _, cid := range allCid {
		_, err := scraper.MergeAllResults[string, int](
			FetchProblemSum(uid, cid),
		)
		//fmt.Println(cRes)
		if err != nil {
			return nil, err
		}

	}
	res["atc_problem_sums"] = len(problemMap)

	return res, nil

}

// 获得 contestUrl
func ScrapeCid(uid string) (map[string]string, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, string](
		FetchContestHistory(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
