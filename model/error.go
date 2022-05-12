package model

import (
	"github.com/FengZhg/go_tools/errs"
)

const (
	errorHTTPResponseStatus = 259001
	errorScrapeTimeout      = 259002
	errorScrapeCallType     = 259003
	errorConfigNotFound     = 259004
)

var (
	HTTPResponseStatusError = errs.NewError(errorHTTPResponseStatus, "HTTP请求相应状态错误")
	ScrapeTimeoutError      = errs.NewError(errorScrapeTimeout, "处理调度超时")
	ScrapeCallTypeError     = errs.NewError(errorScrapeCallType, "scrape类型错误")
	ConfigNotFoundError     = errs.NewError(errorConfigNotFound, "配置文件缺失")
)
