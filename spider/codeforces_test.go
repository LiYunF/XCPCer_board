package spider

import (
	"context"
	"fmt"
	"testing"
)

func TestCodeForces(t *testing.T) {
	keyCheck := [5]string{
		"CodeForces_Practice_PassAmount",
		"CodeForces_Last_Month_Practice_PassAmount",
		"CodeForces_Main_Rating",
		"CodeForces_Main_Max_Rating",
		"CodeForces_Main_Rating_Name",
	}

	cases := []struct {
		Name string
		val  [5]string
	}{
		{"MiracleFaFa",
			[5]string{
				"1310",
				"18",
				"3466",
				"3681",
				"Legendary Grandmaster ",
			},
		},
		{"cityofstar",
			[5]string{
				"349",
				"0",
				"1742",
				"1837",
				"Expert ",
			},
		},
		{"FengZ",
			[5]string{
				"339",
				"9",
				"1761",
				"1922",
				"Expert ",
			},
		},
	}

	for _, c := range cases {
		testPerson, err := GetCFContestPassAmount(context.Background(), c.Name)
		if err != nil {
			t.Errorf("can't get (%v) http", c.Name)
			continue
		}
		for i, e := range testPerson {
			if e.key != keyCheck[i] {
				fmt.Printf(e.key)
				t.Errorf("(%v) :value of key (%v) get err:(%v)", c.Name, keyCheck[i], e.key)
			}
			if e.value != c.val[i] {
				t.Errorf("(%v) :value of value (%v=%v) get err:(%v)", c.Name, keyCheck[i], c.val[i], e.value)
			}

		}
	}

}
