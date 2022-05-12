package scraper

import log "github.com/sirupsen/logrus"

// @Author: Feng
// @Date: 2022/5/12 13:36

//Persist 持久化执行
type Persist func() error

//PersistHandler 持久化处理
type PersistHandler struct {
	cb func(string, interface{}) error
}

func (p *PersistHandler) Do(key string, val interface{}) error {
	return p.cb(key, val)
}

func NewPersistHandler(cb func(string, interface{}) error) *PersistHandler {
	return &PersistHandler{
		cb: cb,
	}
}

var (
	// 持久化任务分配管道
	persistChannel = make(chan Persist)
)

//newPersistProcessor 新持久化处理器
func newPersistProcessor() {
	for persist := range persistChannel {
		err := persist()
		if err != nil {
			log.Errorf("Run Persist Error %v", err)
		}
	}
}

//RegisterPersist 注册持久化处理任务
func RegisterPersist(persists ...Persist) {
	for ind, _ := range persists {
		persistChannel <- persists[ind]
	}
}
