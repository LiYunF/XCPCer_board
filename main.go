package main

import (
	"XCPCer_board/spider/example"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	fmt.Println(example.Scrape("10086"))
}
