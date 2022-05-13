package atcoder

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	conScraper = scraper.NewScraper(
		scraper.WithCallback(conCallback),
		scraper.WithThreads(2),
	)
)

var (
	contestId string
	userId    string
	pageSums  = 1
	subRes    []scraper.KV //保存各个比赛页面的返回信息
)

//conCallback 处理比赛列表的回调函数
func conCallback(c *colly.Collector, res *scraper.Processor) {
	//获取contest页数
	if pageSums == 1 {
		c.OnHTML("ul[class=\"pagination pagination-sm mt-0 mb-1\"]",
			func(element *colly.HTMLElement) {
				ret := element.DOM.Find("li:last-child").First().Text()
				if num, err := strconv.Atoi(ret); err == nil {
					pageSums = num
				}
			},
		)
	}

	//获取比赛id并且进入比赛页面，获取用户提交信息
	c.OnHTML("tbody tr",
		func(element *colly.HTMLElement) {
			//获取比赛id
			cId := strings.Split(element.ChildAttr("td:nth-child(2) a", "href"), "/")[2]
			contestId = cId
			// 进入用户某场比赛提交页面
			ret, err := fetchSubPage(userId, contestId)
			if err != nil && err.Error() != "Not Found" {
				log.Errorf("contestid Fetcher Error %v", err)
			} else {
				subRes = append(subRes, ret...)
			}
		},
	)
}

//getAtCoderPageUrl 获取 userID
func getAtCoderPageUrl(page string) string {
	return "https://atcoder.jp/contests/archive?page=" + page
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchConPage 抓取用户提交所有提交信息
func fetchConPage(uid string) ([]scraper.KV, error) {
	userId = uid
	//进入所有的比赛列表页面
	for num := 1; num <= pageSums; num++ {
		_, err := conScraper.Scrape(getAtCoderPageUrl(strconv.Itoa(num)))
		if err != nil {
			log.Infof("Atcoder Contest Page Error %v", err)
		}
	}
	return subRes, nil
}
