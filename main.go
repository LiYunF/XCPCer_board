package main

import (
	"XCPCer_board/model"
	"XCPCer_board/spider/nowcoder"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	fmt.Println(nowcoder.ScrapeAll(model.TestNowCoderIdWLM))
	//fmt.Println(example.Scrape("10086"))
}
