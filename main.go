package main

import (
	_ "XCPCer_board/config"
	_ "XCPCer_board/dao"
	"XCPCer_board/model"
	"XCPCer_board/spider/atcoder"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	atcoder.Flush(model.TestAtcIdLQY)
}
