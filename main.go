package main

import (
	"XCPCer_board/config"
	"XCPCer_board/db/DBluogu"
	"XCPCer_board/db/mysql"
	"XCPCer_board/model"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {

	//fmt.Println(DBluogu.ScrapeAll(model.TestLuoGuIdLYF))
	config.InitAll()
	//fmt.Println(config.Config)
	mysql.InitDB()
	//fmt.Println(DBluogu.InsertSql(model.TestLuoGuIdLYF))
	fmt.Println(DBluogu.QuerySql(model.TestLuoGuIdLYF))
	//fmt.Println(luogu.ScrapeAll(model.TestLuoGuIdLYF))
}
