package cf

import (
	"XCPCer_board/scraper"
	"fmt"
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
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
		fmt.Println("wtf")

	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}

func InitMsg(uid string) scraper.Results[string] {
	return problemScraper.Scrape(getPersonSubmissionsInit(uid))
}
func GetMsg(uid string, from int, count int) scraper.Results[string] {
	return problemScraper.Scrape(getPersonSubmissions(uid, from, count))
}
