package nowcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
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
	passAmountKey = "nowcoder_pass_amount"

	// 个人练习selector关键字
	passAmountKeyWord = "题已通过"
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
			uid := element.Request.Ctx.Get("uid")
			// 题目通过数量
			num, err := strconv.Atoi(element.DOM.Find(getNowCoderContestBaseFindRule(passAmountKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(getPassAmountKey(uid), num)
		},
	)
}

//getContestPracticeUrl 获取牛客竞赛区个人练习URL
func getContestPracticeUrl(nowCoderId string) string {
	return getContestProfileUrl(nowCoderId) + "/practice-coding"
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
	return practiceScraper.Scrape(func(c *colly.Collector) error {
		ctx := colly.NewContext()
		ctx.Put("uid", uid)
		err := c.Request("GET", getContestPracticeUrl(uid), nil, ctx, nil)
		if err != nil {
			log.Errorf("scraper error %v", err)
			return err
		}
		return nil
	})
}
