package main

import (
	"XCPCer_board/config"
	_ "XCPCer_board/db/Redis"
	_ "XCPCer_board/db/mysql"

	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {


	fmt.Println(config.Config.Database.RedisConf.Host)
	//fmt.Println(luogu.ScrapeAll(model.TestLuoGuIdLYF))


}
