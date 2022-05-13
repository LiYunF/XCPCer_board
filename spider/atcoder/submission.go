package atcoder

import (
	"XCPCer_board/scraper"
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
	conScraper = scraper.NewScraper(
		scraper.WithCallback(mainCallback),
		scraper.WithThreads(2),
	)
)

var (
	contestId string
	userId    string
	pageSums  = 0
	num       = 1
	//subRes       []func(contestId,userId) ([]scraper.KV, error)
)

//conCallback 处理比赛列表的回调函数
func conCallback(c *colly.Collector, res *scraper.Processor) {
	//用goquery
	if pageSums == 0 {
		c.OnHTML("ul[class=\"pagination pagination-sm mt-0 mb-1\"]",
			func(element *colly.HTMLElement) {
				ret := element.DOM.Find("li:last-child").First().Text()
				if num, err := strconv.Atoi(ret); err == nil {
					pageSums = num
				}
			},
		)
	}
	c.OnHTML("tbody tr",
		func(element *colly.HTMLElement) {
			cId := strings.Split(element.ChildAttr("td:nth-child(2) a", "href"), "/")[2]
			contestId = cId
			//ret,err:=fetchSubPage(userId,contestId)
			//if err == nil {append(res,ret...)}
		},
	)
}

//getAtCoderPageUrl 获取 userID
func getAtCoderPageUrl(page string) string {
	return "https://atcoder.jp/contests/archive?page=" + page
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func fetchConPage(uid string) ([]scraper.KV, error) {
	userId = uid
	return mainScraper.Scrape(getAtCoderBaseUrl(uid))
}
