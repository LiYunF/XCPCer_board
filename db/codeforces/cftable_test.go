package codeforces

import "testing"

func TestCreateTable(t *testing.T) {
	err := createTableName()
	if err != nil {
		t.Errorf("wrong create cf table", err)
	}
}
