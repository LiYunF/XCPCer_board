package example

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
)

// @Author: Feng
// @Date: 2022/4/8 17:06

var exampleScraper *scraper.Scraper[int]

func init() {
	exampleScraper = scraper.NewScraper[int](
		scraper.WithCallback(exampleCallback),
	)
}

func exampleCallback(c *colly.Collector, res *scraper.Results[int]) {
	c.OnRequest(func(r *colly.Request) {
		//fmt.Println(r.URL)
		res.Set("Default Callback 1", 1)
		//res.SetError(errs.NewError(0, "Test Error"))
	})
	c.OnHTML("", func(r *colly.HTMLElement) {
		fmt.Println(r.DOM.First().Text())
		res.Set("Default Callback 2", 2)
	})
}

func Scrape(uid string) (map[string]int, error) {
	d, err := exampleScraper.Scrape("https://cn.bing.com")
	if err != nil {
		return nil, err
	}
	return d, nil
}
