package atcoder

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
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
		scraper.WithCallback(mainCallback),
		scraper.WithThreads(2),
	)
)

//mainCallback 处理个人主页的回调函数
func subCallback(c *colly.Collector, res *scraper.Processor) {
	//用goquery
	c.OnHTML("tbody tr",
		func(element *colly.HTMLElement) {
			task := strings.Split(element.DOM.Find("td:nth-child(2)").First().Text(), "")[0]
			score, errSc := strconv.Atoi(element.DOM.Find("td:nth-child(5)").First().Text())
			SMid := strings.Split(element.ChildAttr("td:nth-child(10) a", "href"), "/")[4]
			res.Set(submissionKey+"_"+contestId+"_"+task, submission{userId, SMid, contestId, task, score})
		})
}

//getAtCoderUrl 获取 userID
func getAtCoderUrl(atCoderId string, contestId string) string {
	return "https://atcoder.jp/contests/" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC"
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有
func fetchSubPage(uid string, cid string) ([]scraper.KV, error) {
	return mainScraper.Scrape(getAtCoderUrl(uid, cid))
}
