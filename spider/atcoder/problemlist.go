package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
)

const (
//key

//keyword

)

var (
	problemScraper *scraper.Scraper[int]
	contestId      string
	problemMap     map[string]bool
)

// 初始化
func init() {
	problemScraper = scraper.NewScraper[int](
		scraper.WithCallback(problemCallback),
		scraper.WithThreads[int](2),
	)
}

//处理 acProblem 的页面回调
func problemCallback(c *colly.Collector, res *scraper.Results[int]) {
	//用goquery
	problemMap = make(map[string]bool)

	c.OnHTML("tbody tr", func(element *colly.HTMLElement) {
		ret := element.DOM.Find(fmt.Sprintf("td:nth-child(2)")).First().Text()
		problemMap[ret] = true
		//fmt.Println(len(problemMap))
	})
	res.Set(contestId, len(problemMap))
}

//获取 userID
func getAtCoderUrl(atCoderId string, contestId string) string {
	//fmt.Println("https://atcoder.jp" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC")
	return "https://atcoder.jp" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC"
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有

func FetchProblemSum(uid string, cid string) scraper.Results[int] {
	contestId = cid
	return problemScraper.Scrape(getAtCoderUrl(uid, cid))
}
