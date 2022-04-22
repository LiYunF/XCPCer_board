package cf

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	problemScraper *scraper.Scraper[string]
)

func init() {
	problemScraper = scraper.NewScraper[string](
		scraper.WithCallback(intCallback),
		scraper.WithThreads[string](2),
	)
}

func intCallback(c *colly.Collector, res *scraper.Results[string]) {
	c.OnHTML("#head", func(e *colly.HTMLElement) {
		//fmt.Println(r.DOM.First().Text())

	})
}

func InitMsg(uid string) scraper.Results[string] {
	return problemScraper.Scrape(getPersonSubmissionsInit(uid))
}
func GetMsg(uid string, from int, count int) scraper.Results[string] {
	return problemScraper.Scrape(getPersonSubmissions(uid, from, count))
}

