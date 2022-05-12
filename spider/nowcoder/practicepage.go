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
	// 个人练习页面
	practicePassAmountKey = "NowCoder_Practice_PassAmount"

	// 个人练习selector关键字
	practicePassAmountKeyWord = "题已通过"
)

var (
	practiceScraper *scraper.Scraper[int]
)

// 初始化
func init() {
	practiceScraper = scraper.NewScraper[int](
		scraper.WithCallback(practiceCallback),
	)
}

//practiceCallback 处理牛客个人练习页面的回调函数
func practiceCallback(c *colly.Collector, res *scraper.Processor[int]) {
	//用goquery
	c.OnHTML("html", func(element *colly.HTMLElement) {
		res.Set(practicePassAmountKey, passAmountHandler(element.DOM))
	})
}

//getNowCoderContestProfilePracticeUrl 获取牛客竞赛区个人练习URL
func getNowCoderContestProfilePracticeUrl(nowCoderId string) string {
	return getNowCoderContestProfileBaseUrl(nowCoderId) + "/practice-coding"
}

//getNowCoderContestBaseFindRule 获取牛客竞赛区基础的
func getNowCoderContestBaseFindRule(keyWord string) string {
	return fmt.Sprintf(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix "+
		".my-state-main .my-state-item:contains(%v) .state-num", keyWord)
}

//passAmountHandler 获取竞赛区题目通过数量handler
func passAmountHandler(doc *goquery.Selection) int {
	ret := doc.Find(getNowCoderContestBaseFindRule(practicePassAmountKeyWord)).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//---------------------------------------------------------------------//
// 对外暴露函数:个人练习信息获取
//---------------------------------------------------------------------//

//FetchPractice 抓取个人练习页面的所有
func FetchPractice(uid string) ([]scraper.KV[int], error) {
	return practiceScraper.Scrape(getNowCoderContestProfilePracticeUrl(uid))
}
