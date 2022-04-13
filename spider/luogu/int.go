package luogu

import (
	"XCPCer_board/scraper"
	"encoding/json"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	intScraper *scraper.Scraper[int64]
	jsonData   UserShow
	difficulty [5]int64
)

func init() {
	intScraper = scraper.NewScraper[int64](
		scraper.WithCallback(intCallback),
		scraper.WithThreads[int64](2),
	)
}

func intCallback(c *colly.Collector, res *scraper.Results[int64]) {
	c.OnHTML("head", func(e *colly.HTMLElement) {

		//decoder
		text, _ := url.QueryUnescape(e.DOM.Text())

		//get JsonText
		Data := text[strings.Index(text, "{") : strings.LastIndex(text, "}")+1]
		err := json.Unmarshal([]byte(Data), &jsonData)
		if err != nil {
			log.Println("json Unmarshal error: ", err)
		}
		if jsonData.Code != 200 {
			log.Println("http Response is not 200: ", err)
		}

		//count problem difficulty
		user := jsonData.GetCurrentData().GetUser()
		problem := jsonData.GetCurrentData().GetPassedProblems()

		for i := 0; i < 5; i++ {
			difficulty[i] = 0
		}
		for _, i := range problem {
			q := i.GetDifficulty()
			if q == 0 || q > 7 { //未知题
				difficulty[0]++
			} else if q < 2 { //入门就是简单 q=1
				difficulty[1]++
			} else if q < 4 { //普及-就是基础 q=2,3
				difficulty[2]++
			} else if q < 6 { //普及/提高-,普及+/提高 是进阶 q=4,5
				difficulty[3]++
			} else { //困难
				difficulty[4]++
			}
		}

		//set data
		res.Set(luoGuPersonPassProblemNumber, user.GetPassedProblemCount())
		res.Set(luoGuPersonRanting, user.GetRanking())
		//set data of problem
		res.Set(luoGuUnKnowProblemNumber, difficulty[0])
		res.Set(luoGuSimpleProblemNumber, difficulty[1])
		res.Set(luoGuBasicProblemNumber, difficulty[2])
		res.Set(luoGuElevatedProblemNumber, difficulty[3])
		res.Set(luoGuHardProblemNumber, difficulty[4])
	})
}

func GetStrMsg(uid string) scraper.Results[int64] {
	return intScraper.Scrape(getPersonPractice(uid))
}
