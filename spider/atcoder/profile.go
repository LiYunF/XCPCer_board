package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
)

//---------------------------------------------------------------------//
// atCoder个人信息 //
//---------------------------------------------------------------------//
//  Key

const (
	//key

	RatingKey     = "atc_rating"
	contestSumKey = "atc_contest_sum"

	//keyword

	//RatingKeyword = "Rank"
)

var (
	mainScraper *scraper.Scraper[int]
	contestUrl  []string
)

// 初始化
func init() {
	mainScraper = scraper.NewScraper[int](
		scraper.WithCallback(mainCallback),
		scraper.WithThreads[int](2),
	)
}

//mainCallback 处理个人主页的回调函数
func mainCallback(c *colly.Collector, res *scraper.Results[int]) {
	//用goquery
	c.OnHTML("html", func(element *colly.HTMLElement) {
		res.Set(RatingKey, ratingHandler(element.DOM))
		res.Set(contestSumKey, contestSumHandler(element.DOM))
	})
}

// 获取个人主页URL
func getAtCoderBaseUrl(atCoderId string) string {
	//fmt.Println("test")
	return "https://atcoder.jp/users/" + atCoderId
}

//获取个人rating
func ratingHandler(doc *goquery.Selection) int {
	ret := doc.Find(fmt.Sprintf("body #main-div #main-container .row .col-md-9.col-sm-12 .dl-table.mt-2 " +
		" tr:nth-child(2) td span:first-child")).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//获取比赛场数
func contestSumHandler(doc *goquery.Selection) int {
	ret := doc.Find(fmt.Sprintf("body #main-div #main-container .row .col-md-9.col-sm-12 .dl-table.mt-2 " +
		" tr:nth-child(4) td")).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有
func FetchMainPage(uid string) scraper.Results[int] {
	return mainScraper.Scrape(getAtCoderBaseUrl(uid))
}
