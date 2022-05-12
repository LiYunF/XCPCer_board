package atcoder

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// setUserMsg 插入用户信息
func setUserMsg(ctx *gin.Context, uid string) error {
	res, err := ScrapeAllProfile(uid)
	if err != nil {
		fmt.Println(res)
		//log.Errorf()
	}
	return nil
}
