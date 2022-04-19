package vjudge

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
	c.OnHTML("body", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.First().Text())
		res.Set(vjPersonLast24HoursPassNumber, strToInt(e.DOM, vjPersonLast24HoursPassNumberHandler))
		res.Set(vjPersonLast7DaysPassNumber, strToInt(e.DOM, vjPersonLast7DaysPassNumberHandler))
		res.Set(vjPersonLast30DaysPassNumber, strToInt(e.DOM, vjPersonLast30DaysPassNumberHandler))
		res.Set(vjPersonPassNumber, strToInt(e.DOM, vjPersonPassNumberHandler))
	})
}

func GetIntMsg(uid string) scraper.Results[int] {
	return intScraper.Scrape(getPersonPage(uid))
}
