package codeforcesv2

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCodeforcesApi(t *testing.T) {
	log.SetReportCaller(true)
	userName := "Sariel_snow"

	//get api https://codeforces.com/settings/api
	fc1, err := ScrapeAllWithKey(userName, "26827a9121c91af42b0ef134896f899dc2ecf146", "b7674f19476573d743758bb75a5f548ed6cbfbec")
	if err != nil {
		fmt.Println(err)
	}

	if fc1["154202641"].ContestId == 1668 {
		return
	} else {
		t.Errorf("no")
	}
}
