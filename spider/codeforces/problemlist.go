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
	lastId     int
)

func init() {
	proScraper = scraper.NewScraper[string](
		scraper.WithCallback(proCallback),
		scraper.WithThreads[string](2),
	)
}

func proCallback(c *colly.Collector, res *scraper.Results[string]) {
	c.OnHTML(".datatable table[class=\"status-frame-datatable\"] tbody", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.First().Text())
		fmt.Println("wtf")
		//fmt.Println(e.DOM.First().Text())
		tp := make(map[string]string)
		//minId:=9223372036854775806
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			tp[el.Attr("data-submission-id")] = "1"
		})
		fmt.Println(tp)
	})
}

func GetInitPersonProblemList(uid string, last int) scraper.Results[string] {
	lastId = last
	return proScraper.Scrape(getPersonProblemPage(uid))
}
