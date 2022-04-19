package main

import (
	"XCPCer_board/model"
	"XCPCer_board/spider/vjudge"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {

	//fmt.Println(nowcoder.ScrapeAll(model.TestNowCoderIdLYF))
	fmt.Println(vjudge.ScrapeAll(model.TestVJIdLYF))
}
