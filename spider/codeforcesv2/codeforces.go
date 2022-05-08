package codeforcesv2

import (
	"XCPCer_board/scraper"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	intScraper *scraper.Scraper[Submission]
	jsonData   Status
	difficulty [50]int
)

func init() {
	intScraper = scraper.NewScraper[Submission](
		scraper.WithCallback(intCallback),
		scraper.WithThreads[Submission](1),
	)
}

func intCallback(c *colly.Collector, res *scraper.Results[Submission]) {

	//Json test should use onScraped instead of onHtml
	c.OnScraped(func(r *colly.Response) {

		//get JsonText
		err := json.Unmarshal([]byte(r.Body), &jsonData)
		if err != nil {
			log.Println("json Unmarshal err", err)
		}
		if jsonData.Status != "OK" {
			log.Println("http response is not 200:")
		}

		//count problem difficulty, only if 0 means unknown

		problem := jsonData.GetResult()

		for _, i := range problem {
			if i.Verdict != "OK" {
				continue
			}
			//insert into map
			res.Set(strconv.Itoa(int(i.Id)), *i)
		}

	})

}

// GetStrMsg without key and secret
func GetStrMsg(uid string) scraper.Results[Submission] {
	return intScraper.Scrape(statusWithoutKey(uid))
}

// GetStrMsgWithKey with key and secret
func GetStrMsgWithKey(uid string, key string, secret string) scraper.Results[Submission] {
	return intScraper.Scrape(statusWithKey(uid,
		key,
		secret))

}

func ScrapeAll(uid string) (map[string]Submission, error) {
	// 请求所有并合并所有
	fmt.Println("here i got it")
	res, err := scraper.MergeAllResults[string, Submission](
		GetStrMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ScrapeAllWithKey(uid string, key string, secret string) (map[string]Submission, error) {
	// 请求所有并合并所有
	fmt.Println("here i got it")
	res, err := scraper.MergeAllResults[string, Submission](
		GetStrMsgWithKey(uid, key, secret),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
