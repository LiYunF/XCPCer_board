package codeforces

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	proScraper *scraper.Scraper[string]
)

func init() {
	proScraper = scraper.NewScraper[string](
		scraper.WithCallback(proCallback),
		scraper.WithThreads[string](2),
	)
}

func proCallback(c *colly.Collector, res *scraper.Results[string]) {
	c.OnHTML("body", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.First().Text())
		fmt.Println()
	})
}

func GetInitPersonProblemList(uid string) scraper.Results[string] {
	return proScraper.Scrape(getPersonPage(uid))
}
