package main

import (
	"XCPCer_board/model"
	"XCPCer_board/spider"
	"context"
	"fmt"
	_ "github.com/FengZhg/go_tools/logrus_gin_log"
)

// 主入口函数
func main() {
	//fmt.Println(spider.GetNCContestPassAmount(context.Background(), model.TestNowCoderIdWLM))
	//fmt.Println(spider.GetNCContestRating(context.Background(), model.TestNowCoderIdWLM))
	checkCodeforcesReturn, err := spider.GetCFContestPassAmount(context.Background(), model.TestCodeForcesIdTeacherDu)
	if err != nil {
		fmt.Println("no way")
	}
	fmt.Println(checkCodeforcesReturn)
	for i, e := range checkCodeforcesReturn {
		fmt.Println(i, *e)
	}

}
