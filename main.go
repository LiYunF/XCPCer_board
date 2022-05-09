package main

import (
	_ "XCPCer_board/config"
	_ "XCPCer_board/dao"
	"XCPCer_board/model"
	"XCPCer_board/spider/luogu"
	"context"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	ctx := context.Background()
	fmt.Println(luogu.SetUserMsgToRedis(model.TestLuoGuIdLYF, ctx))
	fmt.Println(luogu.GetUserMsgFromRedis(model.TestLuoGuIdLYF, luogu.UserList[0], ctx))
	fmt.Println(luogu.GetUserAllMsgFromRedis(model.TestLuoGuIdLYF, ctx))

	//fmt.Println(nowcoder.ScrapeAll(model.TestNowCoderIdLYF))
}
