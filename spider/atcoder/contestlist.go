package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

const (
	//key
	contestKey = "atc_contest_id"

	//keyword

)

var (
	contestScraper *scraper.Scraper[string]
	flag           bool
	pageSums       int
	Page           string
	num            = 1
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
	c.OnHTML("tbody tr", func(element *colly.HTMLElement) {

		str := strconv.Itoa(num)
		fmt.Println(str)
		res.Set(contestKey+"_"+str, getAtCoderContestId(element))
		num = num + 1
	})
}

//获取 userID
func getAtCoderPageUrl(page string) string {
	//fmt.Println("https://atcoder.jp/contests/archive?page=" + page)

	return "https://atcoder.jp/contests/archive?page=" + page

}

//获取 contestId
func getAtCoderContestId(e *colly.HTMLElement) string {
	//fmt.Println(e.DOM.Find("td:nth-child(2) a").First().Text())
	link := e.ChildAttr("td:nth-child(2) a", "href")
	link = strings.Split(link, "/")[2]
	//fmt.Println(link)
	return link
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有

func FetchContestPage(page string) scraper.Results[string] {
	Page = page
	return contestScraper.Scrape(getAtCoderPageUrl(page))
}
