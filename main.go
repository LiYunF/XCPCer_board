package main

import (
	_ "XCPCer_board/db/mysql"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {


	fmt.Println(luogu.ScrapeSub(model.TestLuoGuIdLYF))
	//fmt.Println(luogu.ScrapeAll(model.TestLuoGuIdLYF))

}
