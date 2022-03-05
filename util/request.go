package util

import (
	"XCPCer_board/model"
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//SendHTTPGet 发送HTTP请求
func SendHTTPGet(ctx context.Context, url string) (*http.Response, error) {

	// 发送Get请求
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("Send HTTP Get Error url = %v \t err = %v", url, err)
		return nil, err
	}

	// 判断响应状态
	if res.StatusCode != http.StatusOK {
		log.Errorf("HTTP Request Error Url = %v \t Status Code = %v \t Status: %v", url, res.StatusCode, res.Status)
		return nil, model.HTTPResponseStatusError
	}

	return res, err
}
