package db

import (
	"fmt"
	"testing"
)

func TestGetConn(t *testing.T) {
	var tests = []struct {
	}{
		{},
	}
	for _, _ = range tests {
		testname := fmt.Sprintf("MySQL driver")
		t.Run(testname, func(t *testing.T) {
			_, err := GetConn()
			if err != nil {
				t.Errorf("got error, %s", err)
			}
		})
	}
}
