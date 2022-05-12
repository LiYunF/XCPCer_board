package nowcoder

import (
	"XCPCer_board/scraper"
	"fmt"
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
	practiceScraper = scraper.NewScraper(
		scraper.WithCallback(practiceCallback),
	)
)

//practiceCallback 处理牛客个人练习页面的回调函数
func practiceCallback(c *colly.Collector, res *scraper.Processor) {
	//用goquery
	c.OnHTML(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix .my-state-main",
		func(element *colly.HTMLElement) {
			// 题目通过数量
			ret := element.DOM.Find(getNowCoderContestBaseFindRule(practicePassAmountKeyWord)).First().Text()
			if num, err := strconv.Atoi(ret); err == nil {
				res.Set(practicePassAmountKey, num)
			}
		},
	)
}

//getNowCoderContestProfilePracticeUrl 获取牛客竞赛区个人练习URL
func getNowCoderContestProfilePracticeUrl(nowCoderId string) string {
	return getNowCoderContestProfileBaseUrl(nowCoderId) + "/practice-coding"
}

//getNowCoderContestBaseFindRule 获取牛客竞赛区基础的
func getNowCoderContestBaseFindRule(keyWord string) string {
	return fmt.Sprintf(".my-state-item:contains(%v) .state-num", keyWord)
}

//---------------------------------------------------------------------//
// 对外暴露函数:个人练习信息获取
//---------------------------------------------------------------------//

//fetchPractice 抓取个人练习页面的所有
func fetchPractice(uid string) ([]scraper.KV, error) {
	return practiceScraper.Scrape(getNowCoderContestProfilePracticeUrl(uid))
}
