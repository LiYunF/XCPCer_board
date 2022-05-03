package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
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
	SMid   int    //提交编号
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

//处理 acProblem 的页面回调
func problemCallback(c *colly.Collector, res *scraper.Results[submission]) {
	num := 1
	c.OnHTML("tbody tr", func(element *colly.HTMLElement) {
		res.Set(submissionKey+"_"+strconv.Itoa(num), getAtcSubMsg(element))
		num = num + 1
	})

}

func getAtcSubMsg(e *colly.HTMLElement) submission {
	task := fmt.Sprint(e.DOM.Find("td:nth-child(2)").First().Text()[0])
	score, errSc := strconv.Atoi(e.DOM.Find("td:nth-child(5)").First().Text())
	SMid, errSm := strconv.Atoi(e.ChildAttr("td:nth-child(10)", "href"))
	if errSc != nil {
		return submission{userId, SMid, contestId, task, -1}
	}
	if errSm != nil {
		return submission{userId, -1, contestId, task, -1}
	}
	return submission{userId, SMid, contestId, task, score}
}

//获取 userID
func getAtCoderUrl(atCoderId string, contestId string) string {
	//fmt.Println("https://atcoder.jp" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC")
	return "https://atcoder.jp" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC"
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有

func FetchProblemSum(uid string, cid string) scraper.Results[submission] {
	contestId = cid
	userId = uid
	return problemScraper.Scrape(getAtCoderUrl(uid, cid))
}
