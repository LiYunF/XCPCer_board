package nowcoder

import "github.com/imdario/mergo"

// @Author: Feng
// @Date: 2022/4/8 17:09

//ScrapeAll 拉取牛客的所有结果
func ScrapeAll(uid string) (map[string]int, error) {
	res := map[string]int{}
	mainRet, err := FetchMainPage(uid)
	if err != nil {
		return nil, err
	}
	practiceRet, err := FetchPractice(uid)
	if err != nil {
		return nil, err
	}
	err = mergo.Merge(res, mainRet)
	if err != nil {
		return nil, err
	}
	err = mergo.Merge(res, practiceRet)
	if err != nil {
		return nil, err
	}
	return res, nil
}
