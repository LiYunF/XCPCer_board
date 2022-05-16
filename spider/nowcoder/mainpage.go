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
	// 个人主页
	mainRatingKey              = "rating"
	mainRatingRatingKey        = "rating ranking"
	mainAttendContestAmountKey = "attend contest amount"

	// 个人主页selector关键字
	mainRatingKeyWord              = "Rating"
	mainRatingRankingKeyWord       = "Rating排名"
	mainAttendContestAmountKeyWord = "次比赛"
)

var (
	mainScraper = scraper.NewScraper(
		scraper.WithCallback(mainCallback),
		scraper.WithThreads(2),
	)
)

//mainCallback 处理牛客个人主页的回调函数
func mainCallback(c *colly.Collector, res *scraper.Processor) {
	//用goquery
	c.OnHTML(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix .my-state-main",
		func(element *colly.HTMLElement) {
			// rating
			ret := element.DOM.Find(fmt.Sprintf(".my-state-item:contains(%v) .state-num.rate-score5",
				mainRatingKeyWord)).First().Text()
			num, err := strconv.Atoi(ret)
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(mainRatingKey, num)
			// 排名
			ret = element.DOM.Find(getNowCoderContestBaseFindRule(mainRatingRankingKeyWord)).First().Text()
			num, err = strconv.Atoi(ret)
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(mainRatingRatingKey, num)
			// 过题数
			ret = element.DOM.Find(getNowCoderContestBaseFindRule(mainAttendContestAmountKeyWord)).First().Text()
			num, err = strconv.Atoi(ret)
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(mainAttendContestAmountKey, num)
		},
	)

}

//getNowCoderContestProfileBaseUrl 获取牛客竞赛区个人主页URL
func getNowCoderContestProfileBaseUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func fetchMainPage(uid string) ([]scraper.KV, error) {
	return mainScraper.Scrape(getNowCoderContestProfileBaseUrl(uid))
}
