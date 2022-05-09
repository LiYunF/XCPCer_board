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

var (
	problemScraper *scraper.Scraper[submission]
	contestId      string
	userId         string
	flag           string
)

// submission 信息
type submission struct {
	userid string //用户名
	SMid   string //提交编号
	CTid   string //比赛编号
	task   string //题目序号
	score  int    //题目难度
}

// 初始化
func init() {
	problemScraper = scraper.NewScraper[submission](
		scraper.WithCallback(problemCallback),
		scraper.WithThreads[submission](2),
	)
}

//problemCallback 处理 acProblem 的页面回调
func problemCallback(c *colly.Collector, res *scraper.Results[submission]) {
	c.OnHTML("tbody tr", func(element *colly.HTMLElement) {
		tmp := getAtcSubMsg(element)
		res.Set(submissionKey+"_"+contestId+"_"+tmp.task, tmp)
	})
}

// getAtcSubMsg 获取每条submission信息
func getAtcSubMsg(e *colly.HTMLElement) submission {

	task := strings.Split(e.DOM.Find("td:nth-child(2)").First().Text(), "")[0]

	score, errSc := strconv.Atoi(e.DOM.Find("td:nth-child(5)").First().Text())

	SMid := strings.Split(e.ChildAttr("td:nth-child(10) a", "href"), "/")[4]

	if errSc != nil {
		return submission{userId, "-1", contestId, task, -1}
	}

	return submission{userId, SMid, contestId, task, score}
}

//getAtCoderUrl 获取 userID
func getAtCoderUrl(atCoderId string, contestId string) string {
	return "https://atcoder.jp/contests/" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC"
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchProblemSum 抓取提交页面所有
func FetchProblemSum(uid string, cid string) scraper.Results[submission] {
	contestId = cid
	userId = uid
	return problemScraper.Scrape(getAtCoderUrl(uid, cid))
}
