package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"strconv"
)

//ScrapeAllProfile 拉取个人主页信息
func ScrapeAllProfile(uid string) (map[string]int, error) {
	res, err := scraper.MergeAllResults[string, int](
		FetchMainPage(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//ScrapeSubmission 拉取所有submission信息
func ScrapeSubmission(uid string) (map[string]submission, error) {
	resCid, errC := ScrapeCid()
	if errC != nil {
		return nil, errC
	}

	//遍历contest
	var res []scraper.Results[submission]
	for _, id := range resCid {
		//特判无权限比赛
		if id == "asprocon8" {
			continue
		}
		res = append(res, FetchProblemSum(uid, id))
	}
	//fmt.Println(len(res))
	ans, err := scraper.MergeAllResults[string, submission](res...)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(ans))
	return ans, err
}

// ScrapeCid 获得contestId
func ScrapeCid() (map[string]string, error) {
	pageSums = 1
	var res []scraper.Results[string]

	// 访问 contestPage 的页面
	for pageNum := 1; pageNum <= pageSums; pageNum++ {
		pNum := strconv.Itoa(pageNum)
		res = append(res, FetchContestPage(pNum))
	}
	ans, err := scraper.MergeAllResults[string, string](res...)
	if err != nil {
		return nil, err
	}
	return ans, nil
}
