package atcoder

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

const (
	//key
	submissionKey = "submission_msg"
	//keyword

)

var (
	problemScraper *scraper.Scraper[submission]
	contestId      string
	userId         string
)

// shubmission 信息
type submission struct {
	userid string
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
	num := 1
	c.OnHTML("tbody tr", func(element *colly.HTMLElement) {
		res.Set(submissionKey+"_"+contestId+"_"+strconv.Itoa(num), getAtcSubMsg(element))
		num = num + 1
	})
}

// getAtcSubMsg 获取每条submission信息
func getAtcSubMsg(e *colly.HTMLElement) submission {

	task := e.DOM.Find("td:nth-child(2)").First().Text()
	task = strings.Split(task, "")[0]

	score, errSc := strconv.Atoi(e.DOM.Find("td:nth-child(5)").First().Text())

	SMid := e.ChildAttr("td:nth-child(10) a", "href")
	SMid = strings.Split(SMid, "/")[4]
	//fmt.Println(SMid)
	if errSc != nil {
		return submission{userId, "-1", contestId, task, -1}
	}

	return submission{userId, SMid, contestId, task, score}
}

//getAtCoderUrl 获取 userID
func getAtCoderUrl(atCoderId string, contestId string) string {
	//fmt.Println("https://atcoder.jp" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC")
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
