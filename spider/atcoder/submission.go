package atcoder

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	//key
	submissionKey = "submission"
	//keyword
)

// submission 信息
type submission struct {
	userid string //用户名
	SMid   string //提交编号
	CTid   string //比赛编号
	task   string //题目序号
	score  int    //题目难度
}

var (
	subScraper = scraper.NewScraper(
		scraper.WithCallback(subCallback),
		scraper.WithThreads(2),
	)
)

//subCallback 处理用户比赛提交页面的回调函数
func subCallback(c *colly.Collector, res *scraper.Processor) {
	//获取提交信息
	c.OnHTML("tbody tr",
		func(element *colly.HTMLElement) {
			//题目序号
			task := strings.Split(element.DOM.Find("td:nth-child(2)").First().Text(), "")[0]
			//题目难度
			score, errSc := strconv.Atoi(element.DOM.Find("td:nth-child(5)").First().Text())
			//提交编号
			SMid := strings.Split(element.ChildAttr("td:nth-child(10) a", "href"), "/")[4]
			if errSc != nil {
				log.Errorf("subpage Fetcher Error %v", errSc)
			}
			res.Set(submissionKey, submission{userId, SMid, contestId, task, score})
		})
}

//getAtCoderUrl 获取用户提交页面链接
func getAtCoderUrl(atCoderId string, contestId string) string {
	return "https://atcoder.jp/contests/" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC"
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchSubPage 抓取用户某场比赛所有提交信息
func fetchSubPage(uid string, cid string) ([]scraper.KV, error) {
	return subScraper.Scrape(getAtCoderUrl(uid, cid))
}
