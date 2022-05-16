package vjudge

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	intScraper = scraper.NewScraper(
		scraper.WithCallback(intCallback),
		scraper.WithThreads(2),
	)
)

func intCallback(c *colly.Collector, res *scraper.Processor) {
	c.OnHTML("body", func(e *colly.HTMLElement) {
		//debug:
		//fmt.Println(e.DOM.First().Text())

		//get last 24 Hours pass problem Number
		retStr := e.DOM.Find(fmt.Sprintf(".container a[title=\"New solved in last 24 hours\"]")).First().Text()
		num, err := strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err,the return is %v \n the err is :%v ", retStr, err)
		}
		res.Set(last24HoursNumber, num)

		//get last 7 Days pass problem Number
		retStr = e.DOM.Find(fmt.Sprintf(".container a[title=\"New solved in last 7 days\"]")).First().Text()
		num, err = strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err,the return is %v \n the err is :%v ", retStr, err)
		}
		res.Set(last7DaysNumber, num)

		//get last 30 Days pass problem Number
		retStr = e.DOM.Find(fmt.Sprintf(".container a[title=\"New solved in last 30 days\"]")).First().Text()
		num, err = strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err,the return is %v \n the err is :%v ", retStr, err)
		}
		res.Set(last30DaysNumber, num)

		//get total pass problem Number
		retStr = e.DOM.Find(fmt.Sprintf(".container a[title=\"Overall solved\"]")).First().Text()
		num, err = strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err,the return is %v \n the err is :%v ", retStr, err)
		}
		res.Set(totalNumber, num)
	})
}

//////////////////////对外暴露函数////////////////////////

//GetUserMsg 获取用户信息
func GetUserMsg(uid string) ([]scraper.KV, error) {
	return intScraper.Scrape(getPersonPage(uid))
}
