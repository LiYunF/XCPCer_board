package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
)

//ScrapeAll 拉取个人主页
func ScrapeAll(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		FetchMainPage(uid),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//合并submission
func ScrapeSubmission(uid string) (map[string]submission, error) {

	resCid, errC := ScrapeCid()

	if errC != nil {
		return nil, errC
	}

	//resCid := map[string]string{"1": "abc248"}

	var res map[string]submission
	res = make(map[string]submission)

	for _, id := range resCid {
		if id == "asprocon8" {
			continue
		}
		resCon, errCon := scraper.MergeAllResults[string, submission](
			FetchProblemSum(uid, id),
		)
		if errCon != nil {
			return nil, errCon
		}
		for k, v := range resCon {
			res[k] = v
		}
	}
	fmt.Println(len(res))
	return res, nil

}

// 获得 contestUrl
func ScrapeCid() (map[string]string, error) {
	// 请求所有并合并所有

	res, err := scraper.MergeAllResults[string, string](
		FetchContestPage("1"),
		FetchContestPage("2"),
		FetchContestPage("3"),
		FetchContestPage("4"),
		FetchContestPage("5"),
		FetchContestPage("6"),
		FetchContestPage("7"),
		FetchContestPage("8"),
		FetchContestPage("9"),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
