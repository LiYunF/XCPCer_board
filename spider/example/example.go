package example

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
)

// @Author: Feng
// @Date: 2022/4/8 17:06

func scraperExample() {
	s, err := scraper.NewScraper[int](scraper.WithCallback(func(c *colly.Collector, ch chan int) {
		c.OnScraped(func(res *colly.Response) {
			fmt.Println(res.Request.URL)
			fmt.Println(string(res.Body))
			ch <- len(res.Body)
		})
	}))
	d, err := s.Scrape("https://cn.bing.com")
	fmt.Println(d, err)
}
