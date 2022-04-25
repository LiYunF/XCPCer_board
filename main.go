package main

import (
	"XCPCer_board/config"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {

	//fmt.Println(nowcoder.ScrapeAll(model.TestNowCoderIdLYF))
	//fmt.Println(vjudge.ScrapeAll(model.TestVJIdLYF))
	//fmt.Println(codeforces.GetInitPersonProblemList(model.TestCodeForcesIdLYF, 10000))
	q := config.GetDBMsg()
	fmt.Println(q.CF)
}
