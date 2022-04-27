package atcoder

import (
	"XCPCer_board/model"
	"testing"
)

var atcTestData = map[string]int{
	"atc_contest_sum": 6,
	"atc_rating":      785,
}

func dataCheck(atcData map[string]int) bool {
	for key, val := range atcTestData {
		if _, data := atcData[key]; data == false || atcData[key] != val {
			return true
		}
	}
	return false
}

func TestAtc(t *testing.T) {
	atcRes, err := ScrapeAll(model.TestAtcIdLQY)
	if err != nil {
		t.Errorf("Error of atcoder is : %v", err)
	}
	if dataCheck(atcRes) {
		t.Errorf("Error of atcoder \n ans is %v \n but output is %v", atcTestData, atcRes)
	}
}
