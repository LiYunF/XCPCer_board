package codeforces

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	strScraper *scraper.Scraper[string]
)

func init() {
	strScraper = scraper.NewScraper[string](
		scraper.WithCallback(strCallback),
		scraper.WithThreads[string](2),
	)
}

func strCallback(c *colly.Collector, res *scraper.Results[string]) {
	c.OnHTML("#body", func(e *colly.HTMLElement) {
		//fmt.Println(r.DOM.First().Text())
		res.Set(ratingNameKey, ratingNameHandler(e.DOM))
	})
}

//GetStrMsg 对外暴露函数，获取str信息
func GetStrMsg(uid string) scraper.Results[string] {
	return strScraper.Scrape(getPersonPage(uid))
}
