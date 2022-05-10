package main

import (
	_ "XCPCer_board/config"
	_ "XCPCer_board/dao"
	"XCPCer_board/model"
	"XCPCer_board/spider/vjudge"
	"context"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	ctx := context.Background()
	fmt.Println(vjudge.SetUserMsgToRedis(model.TestVJIdLYF, ctx))
	fmt.Println(vjudge.GetUserMsgFromRedis(model.TestVJIdLYF, vjudge.UserList[0], ctx))
	fmt.Println(vjudge.GetUserAllMsgFromRedis(model.TestVJIdLYF, ctx))
	//fmt.Println(nowcoder.ScrapeAll(model.TestNowCoderIdLYF))
}
