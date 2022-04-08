package example

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/8 17:06

var exampleScraper *scraper.Scraper[int]

func init() {
	e, err := scraper.NewScraper[int](
		scraper.WithCallback(exampleCallback),
	)
	if err != nil {
		panic(err)
	}
	exampleScraper = e
}

func exampleCallback(c *colly.Collector, ch chan int) {
	c.OnScraped(func(res *colly.Response) {
		fmt.Println(res.Request.URL)
		fmt.Println(string(res.Body))
		ch <- len(res.Body)
	})
}

func Scrape(uid string) (int, error) {
	d, err := exampleScraper.Scrape("https://cn.bing.com")
	if err != nil {
		log.Errorf("Example Error %v", err)
		return 0, err
	}
	return d, nil
}
