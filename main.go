package main

import (
	"XCPCer_board/spider/luogu"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	//fmt.Println(nowcoder.ScrapeAll(model.TestNowCoderIdWLM))
	//fmt.Println(example.Scrape("10086"))
	//fmt.Println(codeforces.ScrapeAll("MiracleFaFa"))
	//fmt.Println(codeforces.ScrapeInt("MiracleFaFa"))
	//fmt.Println(codeforces.ScrapeStr("MiracleFaFa"))
	fmt.Println(luogu.ScrapeAll("241858"))
	fmt.Println(luogu.ScrapeAll("241858"))
	fmt.Println(luogu.ScrapeAll("241858"))

}
