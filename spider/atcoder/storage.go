package atcoder

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// setUserMsg 插入用户信息
func setUserMsg(ctx *gin.Context, uid string) error {
	res, err := ScrapeAllProfile(uid)
	if err != nil {
		log.Errorf()
	}
}
