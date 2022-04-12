package codeforces

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	intScraper *scraper.Scraper[int]
)

func init() {
	intScraper = scraper.NewScraper[int](
		scraper.WithCallback(intCallback),
		scraper.WithThreads[int](2),
	)
}

func intCallback(c *colly.Collector, res *scraper.Results[int]) {
	c.OnHTML("#body", func(e *colly.HTMLElement) {
		//fmt.Println(r.DOM.First().Text())
		res.Set(codeforcesPracticePassAmountKey, strToInt(e.DOM, cfPracticePassAmountHandler))
		res.Set(codeforcesLastMonthPracticePassAmount, strToInt(e.DOM, cfPracticePassLastMonthAmountHandler))
		res.Set(codeforcesMainRatingKey, strToInt(e.DOM, cfMainRatingHandler))
		res.Set(codeforcesMaxMainRatingKey, strToInt(e.DOM, cfMaxMainRatingHandler))
	})
}

func GetIntMsg(uid string) scraper.Results[int] {
	return intScraper.Scrape(getPersonPage(uid))
}
