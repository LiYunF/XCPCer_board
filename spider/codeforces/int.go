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
		res.Set(problemPassAmountKey, strToInt(e.DOM, problemPassAmountHandler))
		res.Set(lastMonthPassAmount, strToInt(e.DOM, lastMonthAmountHandler))
		res.Set(ratingKey, strToInt(e.DOM, ratingHandler))
		res.Set(maxRatingKey, strToInt(e.DOM, maxRatingHandler))
	})
}

//GetIntMsg 对外暴露函数，获取int信息
func GetIntMsg(uid string) scraper.Results[int] {
	return intScraper.Scrape(getPersonPage(uid))
}
