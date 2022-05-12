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
	pageSums = 0
	num      = 1
)

//mainCallback 处理牛客个人主页的回调函数
func subCallback(c *colly.Collector, res *scraper.Processor) {
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
			cId := element.ChildAttr("td:nth-child(2) a", "href")
			cId = strings.Split(cId, "/")[2]
			
		},
	)

}

//getAtCoderPageUrl 获取 userID
func getAtCoderPageUrl(page string) string {
	return "https://atcoder.jp/contests/archive?page=" + page
}
