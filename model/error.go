package model

import (
	"github.com/FengZhg/go_tools/errs"
)

const (
	errorHTTPResponseStatus = 259001
	errorScrapeTimeout      = 259002
)

var (
	HTTPResponseStatusError = errs.NewError(errorHTTPResponseStatus, "HTTP请求相应状态错误")
	ScrapeTimeoutError      = errs.NewError(errorScrapeTimeout, "处理调度超时")
)
