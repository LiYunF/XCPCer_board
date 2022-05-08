package dao

import (
	"XCPCer_board/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var DbClient *sql.DB

const mysqlDriver = "mysql"

//初始化数据库的函数
func init() {
	// 判断是否存在配置
	mysqlConfig, ok := config.Conf.Storages[mysqlDriver]
	if !ok {
		panic(fmt.Errorf("lack of mysql Conf"))
	}
	// 初始化连接
	var err error
	DbClient, err = sql.Open(mysqlDriver, fmt.Sprintf("%v:%v@tcp(%v)/", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host))
	if err != nil {
		log.Errorf("Open Sql Error: %v", err)
		panic(err)
	}
}
