package nowcoder

import (
	"XCPCer_board/model"
	"testing"
)

// @Author: Feng
// @Date: 2022/5/12 17:40

func BenchmarkFlush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Flush(model.TestNowCoderIdLYF)
	}
}
