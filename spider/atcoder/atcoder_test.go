package atcoder

import (
	"XCPCer_board/model"
	"testing"
)

var atcTestData = map[string]int{
	"atc_contest_sum": 7,
	"atc_rank":        15279,
	"atc_rating":      837,
}

func TestAtc(t *testing.T) {
	atcRes, err := ScrapeAllProfile(model.TestAtcIdLQY)

	if err != nil {
		t.Errorf("Error of atcoder is : %v", err)
	}

	flag := 0
	for key, val := range atcTestData {
		if _, data := atcRes[key]; data == false || atcRes[key] != val {
			flag = 1
		}
	}
	if flag == 1 {
		t.Errorf("Error of atcoder \n ans is %v \n but output is %v", atcTestData, atcRes)
	}
}
