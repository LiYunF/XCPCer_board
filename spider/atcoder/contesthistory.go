package atcoder

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	"strconv"
)

const (
	//key
	contestKey = "atc_contest_id"

	//keyword

)

var (
	contestScraper *scraper.Scraper[string]
)

// 初始化
func init() {
	contestScraper = scraper.NewScraper[string](
		scraper.WithCallback(contestCallback),
		scraper.WithThreads[string](2),
	)
}

//处理 contestHistory 的页面回调
func contestCallback(c *colly.Collector, res *scraper.Results[string]) {
	//用goquery
	num := 1

	c.OnHTML("tr", func(element *colly.HTMLElement) {
		str := strconv.Itoa(num)
		res.Set(contestKey+"_"+str, getAtCoderContestId(element))
		num = num + 1
	})
}

//获取 userID
func getAtCoderHistoryUrl(page string) string {
	return "https://atcoder.jp/contests/archive?page=" + page
}

//获取 contestId
func getAtCoderContestId(e *colly.HTMLElement) string {
	link := e.ChildAttr("a:first-child", "href")
	//fmt.Println(link)
	return link
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有

func FetchContestpage(page int) scraper.Results[[]string] {
	return contestScraper.Scrape(getAtCoderHistoryUrl(page))
}
