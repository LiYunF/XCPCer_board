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

	//resCid := map[string]string{"1": "abc248"}

	//fmt.Println(resCid)

	var res map[string]submission
	res = make(map[string]submission)

	//遍历contest
	for _, id := range resCid {
		//特判无权限比赛
		if id == "asprocon8" {
			continue
		}
		resCon, errCon := scraper.MergeAllResults[string, submission](
			FetchProblemSum(uid, id),
		)
		if errCon != nil {
			return nil, errCon
		}

		//合并map
		for k, v := range resCon {
			res[k] = v
		}
	}
	fmt.Println(len(res))
	return res, nil

}

// ScrapeCid 获得contestId
func ScrapeCid() (map[string]string, error) {
	// 请求所有并合并所有

	pageSums = 1

	var res map[string]string
	res = make(map[string]string)

	// 访问 contestPage 的页面
	for pageNum := 1; pageNum <= pageSums; pageNum++ {
		pNum := strconv.Itoa(pageNum)
		ans, err := scraper.MergeAllResults[string, string](
			FetchContestPage(pNum),
		)
		if err != nil {
			return nil, err
		}
		for k, v := range ans {
			res[k] = v
		}
	}
	return res, nil
}
