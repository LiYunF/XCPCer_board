package nowcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
)

// @Author: Feng
// @Date: 2022/4/11 16:17

//-------------------------------------------------------------------------------------------//
// 基础方法
//-------------------------------------------------------------------------------------------//

// 牛客finder存储Key
const (
	// 个人主页
	mainRatingKey              = "NowCoder_Main_Rating"
	mainRatingRatingKey        = "NowCoder_Main_RatingRanking"
	mainAttendContestAmountKey = "NowCoder_Main_AttendContestAmount"

	// 个人主页selector关键字
	mainRatingKeyWord              = "Rating"
	mainRatingRankingKeyWord       = "Rating排名"
	mainAttendContestAmountKeyWord = "次比赛"
)

var (
	mainScraper *scraper.Scraper[int]
)

// 初始化
func init() {
	mainScraper = scraper.NewScraper[int](
		scraper.WithCallback(mainCallback),
		scraper.WithThreads[int](2),
	)
}

//mainCallback 处理牛客个人主页的回调函数
func mainCallback(c *colly.Collector, res *scraper.Processor[int]) {
	//用goquery
	c.OnHTML("html", func(element *colly.HTMLElement) {
		res.Set(mainRatingKey, ratingHandler(element.DOM))
		res.Set(mainRatingRatingKey, ratingRankingHandler(element.DOM))
		res.Set(mainAttendContestAmountKey, attendContestAmountHandler(element.DOM))
	})

}

//getNowCoderContestProfileBaseUrl 获取牛客竞赛区个人主页URL
func getNowCoderContestProfileBaseUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//ratingHandler 获取竞赛区Rating handler (需要的条件多一个比较特殊)
func ratingHandler(doc *goquery.Selection) int {
	ret := doc.Find(fmt.Sprintf(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix "+
		".my-state-main .my-state-item:contains(%v) .state-num.rate-score5", mainRatingKeyWord)).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//ratingRankingHandler 获取竞赛区Rating排名handler
func ratingRankingHandler(doc *goquery.Selection) int {
	ret := doc.Find(getNowCoderContestBaseFindRule(mainRatingRankingKeyWord)).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//attendContestAmountHandler 获取竞赛区Rating排名handler
func attendContestAmountHandler(doc *goquery.Selection) int {
	ret := doc.Find(getNowCoderContestBaseFindRule(mainAttendContestAmountKeyWord)).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有
func FetchMainPage(uid string) ([]scraper.KV[int], error) {
	return mainScraper.Scrape(getNowCoderContestProfileBaseUrl(uid))
}
