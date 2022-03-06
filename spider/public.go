package spider

import (
	"XCPCer_board/util"
	"context"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"sync"
)

// 处理goquery匹配的回调函数
type goQueryFindHandler func(*goquery.Document) string

type goQueryFinder struct {
	findKey     string
	findHandler goQueryFindHandler
}

// 包裹goquery匹配的返回值
type goQueryFinderReturn struct {
	key, value string
}

// 由于需要手动关闭http client的body进行连接复用 使用回调函数的方式，保证可以关闭读响应的流
//doHTTPGetAndGoQuery 进行http请求和html解析
func doHTTPGetAndGoQuery(ctx context.Context, url string, finders ...goQueryFinder) ([]*goQueryFinderReturn, error) {

	// 请求阶段，并完成请求相应状态错误判断
	res, err := util.SendHTTPGet(ctx, url)
	if err != nil {
		log.Errorf("HTTP Get Error err = %v", err)
		return nil, err
	}

	// 关闭io读，方便连接复用
	defer res.Body.Close()

	//解析html阶段
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Errorf("GoQuery Pharse HTML Error err = %v", err)
		return nil, err
	}

	// 构造返回体，保证并发安全
	rets, wg := make([]*goQueryFinderReturn, len(finders)), sync.WaitGroup{}

	// 批量执行回调函数
	for ind, finder := range finders {
		wg.Add(1)
		// 并发执行回调函数
		go func(idx int, f goQueryFinder) {
			defer wg.Done()
			rets[idx] = &goQueryFinderReturn{
				key:   finder.findKey,
				value: f.findHandler(doc),
			}
		}(ind, finder)
	}

	return rets, nil
}
